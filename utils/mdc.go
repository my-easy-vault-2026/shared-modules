package utils

import (
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	// "github.com/petermattis/goid"

	"gitlab.com/jjjj.sinrui2023/goid"
)

func GoroutineID() string {
	// return strconv.FormatInt(goid.Get(), 10)
	return strconv.FormatUint(goid.Get(), 10)
}

var mdc = cache.New(10*time.Minute, 10*time.Minute)
var NonExpiringMDC = cache.New(0, 0)

func SetMDCValue(key string, value string) {

	mdc.SetDefault(GoroutineID()+key, value)
}

func SetNonExpiringMDCValue(key string, value string) {

	NonExpiringMDC.SetDefault(GoroutineID()+key, value)
}

func GetMDCValue(key string) (string, bool) {

	v, ret := mdc.Get(GoroutineID() + key)

	if ret {
		return v.(string), true
	} else {
		v, ret = NonExpiringMDC.Get(GoroutineID() + key)
		if ret {
			return v.(string), true
		}
		return "", false
	}
}

func DeleteMDCValue(key string) {
	mdc.Delete(GoroutineID() + key)
}

func DeleteNonExpiringMDCValue(key string) {
	NonExpiringMDC.Delete(GoroutineID() + key)
}
