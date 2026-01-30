package utils

import "testing"

func TestGoogleAuth(t *testing.T) {

	secret := NewGoogleAuth().GetSecret()

	t.Log("secret = ", secret)

}

func TestVerifyGoogleAuth(t *testing.T) {

	ret := NewGoogleAuth().VerifyCode("HNZFWGPMOILHSBWP", "455947")

	t.Log("ret=", ret)

}
