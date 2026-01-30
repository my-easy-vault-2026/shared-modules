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
					select {
					case l.errChan <- err:
						cancel()
						return
					case <-time.After(100 * time.Microsecond):
						logger.Warn("setnx err chan block")
						cancel()
						return
					}
				}
				if ret.Val() {
					select {
					case l.lockChan <- struct{}{}:
						cancel()
						return
					case <-time.After(100 * time.Microsecond):
						logger.Warn("setnx res chan block")
						cancel()
						return
					}
				}
			}
			time.Sleep(100 * time.Microsecond)
		}
	}(cctx)
	defer cancel()

	select {
	case err := <-l.errChan:
		if err != nil {
			return err
		}
	case <-l.lockChan:
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
