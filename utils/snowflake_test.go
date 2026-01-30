package utils

import (
	"testing"
)

func TestSnowFlake(t *testing.T) {

	snowFlake0 := NewSnowFlake(1, 0)
	t.Log(snowFlake0.Generate())

	snowFlake1 := NewSnowFlake(1, 1)
	t.Log(snowFlake1.Generate())
}

func TestSnowFlakeID(t *testing.T) {
	t.Log(SnowFlakeUserID.Generate())
}
