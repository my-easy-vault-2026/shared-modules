package utils

import "testing"

func TestMDC(t *testing.T) {

	SetMDCValue("a", "bbbb")
	t.Log(GoroutineID())
	t.Log(GetMDCValue("a"))
}
