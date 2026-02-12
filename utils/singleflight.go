package utils

import (
	"sync"

	"github.com/my-easy-vault-2026/shared-modules/common"

	"golang.org/x/sync/singleflight"
)

var (
	SingleFlightGroup map[common.SingleFlightGroup]*singleflight.Group
	singleFlightLock  sync.RWMutex
	singleFlightOnce  sync.Once
)

func RegisterSingleFlight(group common.SingleFlightGroup) {
	singleFlightOnce.Do(func() {
		SingleFlightGroup = make(map[common.SingleFlightGroup]*singleflight.Group)
	})

	ok := func() bool {
		singleFlightLock.RLock()
		defer singleFlightLock.RUnlock()
		_, ok := SingleFlightGroup[group]
		return ok
	}()

	if !ok {
		singleFlightLock.Lock()
		defer singleFlightLock.Unlock()
		_, ok := SingleFlightGroup[group]
		if !ok {
			SingleFlightGroup[group] = &singleflight.Group{}
		}
	}
}

func GetSingleFlight(group common.SingleFlightGroup) *singleflight.Group {
	singleFlightOnce.Do(func() {
		SingleFlightGroup = make(map[common.SingleFlightGroup]*singleflight.Group)
	})

	sfg, ok := func() (*singleflight.Group, bool) {
		singleFlightLock.RLock()
		defer singleFlightLock.RUnlock()
		sfg, ok := SingleFlightGroup[group]
		return sfg, ok
	}()
	if !ok {
		singleFlightLock.Lock()
		defer singleFlightLock.Unlock()
		sfg, ok := SingleFlightGroup[group]
		if !ok {
			sfg = &singleflight.Group{}
			SingleFlightGroup[group] = sfg
		}
		return sfg
	}
	return sfg
}
