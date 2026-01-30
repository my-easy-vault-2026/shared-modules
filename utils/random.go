package utils

import (
	"math/rand"
	"time"
)

func RandString(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func RandInt(min, max int) int {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	if min > max {
		min, max = max, min
	}
	return seededRand.Intn(max-min+1) + min
}

// Float64 returns a random float64 between 0.0 and 1.0.
func RandFloat() float64 {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	return seededRand.Float64()
}

func RandShuffle[T any](slice []T) {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	seededRand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func RandChoice[T any](slice []T) T {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	return slice[seededRand.Intn(len(slice))]
}
