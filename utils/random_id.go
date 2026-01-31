package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomIDWithPrefix(prefix uint64) (uint64, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	generated := rnd.Uint64()
	len := len(strconv.FormatUint(prefix, 10))
	return strconv.ParseUint(strconv.FormatUint(prefix, 10)+strconv.FormatUint(generated, 10)[:19-len], 10, 64)
}

func RandomID() uint64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	generated := rnd.Uint64()
	return generated
}
