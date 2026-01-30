package utils

import (
	"context"
	"sync"
	"time"

	"shared-modules/common"
	"shared-modules/logger"
)

type Locker struct {
	key      string
	locked   bool
	rwLock   sync.RWMutex
	lockChan chan struct{}
	errChan  chan error
}

func NewLocker() *Locker {
	return &Locker{
		lockChan: make(chan struct{}, 0),
		errChan:  make(chan error, 0),
	}
}

func (l *Locker) WaitLock(ctx context.Context, key string, duration time.Duration, wait time.Duration) error {

	l.key = key
	l.rwLock.Lock()
	defer l.rwLock.Unlock()

	reqID, _ := GetMDCValue("reqId")
	cctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(wait),
	)

	go func(cctx context.Context) {
		SetMDCValue("reqId", reqID)
		for {
			if !l.locked {
				ret := RDB.SetNX(cctx, key, time.Now().Unix(), duration)
				if err := ret.Err(); err != nil {
					l.errChan <- err
					cancel()
					return
				}
				if ret.Val() {
					l.lockChan <- struct{}{}
					cancel()
					return
				}
			}
			time.Sleep(Config.System.LockPollingMicroseconds * time.Microsecond)
		}
	}(cctx)
	defer cancel()

	returnByChan := false
	defer func() {
		if !returnByChan {
			go func() {
				SetMDCValue("reqId", reqID)
				select {
				case <-l.errChan:
					return
				case <-l.lockChan:
					RDB.Del(context.Background(), key)
					return
				case <-time.After(duration):
					logger.Warnf("wait too long for leftover lock [%s]", key)
					return
				}
			}()
		}
	}()

	select {
	case err := <-l.errChan:
		returnByChan = true
		if err != nil {
			return err
		}
	case <-l.lockChan:
		returnByChan = true
		return nil
	case <-cctx.Done():
		switch cctx.Err() {
		case context.DeadlineExceeded, context.Canceled:
			return NewBusinessError(ctx, common.CODE_SYSTEM_BUSY)
		}
		return NewBusinessError(ctx, common.CODE_SYSTEM_BUSY)
	case <-time.After(wait):
		logger.Warnf("wait lock timeout [%s]", key)
		return NewBusinessError(ctx, common.CODE_SYSTEM_BUSY)
	}

	return NewBusinessError(ctx, common.CODE_UNKOWN_ERROR)
}

func (l *Locker) Lock(ctx context.Context, key string, duration time.Duration) bool {

	l.key = key
	l.rwLock.Lock()
	defer l.rwLock.Unlock()

	ret := RDB.SetNX(ctx, key, time.Now().Unix(), duration)

	if ret.Err() != nil {
		return false
	}

	return ret.Val()

}

func (l *Locker) UnLock(ctx context.Context) error {

	l.rwLock.Lock()
	defer l.rwLock.Unlock()

	ret := RDB.Del(ctx, l.key)

	if ret.Err() != nil {
		return ret.Err()
	}
	return nil
}
