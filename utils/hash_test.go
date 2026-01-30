package utils

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {

	t.Log(Sha256String("123456"))
	t.Log(HmacSha256(Sha256String("123456"), "3$WTm*#cXtJY"))
}

func TestGeneratePassword(t *testing.T) {
	salt, err := GenerateSalt(8)
	if err != nil {
		fmt.Printf("generate salt failed, %v", err)
	}

	hashedPassword, err := BcryptHash("111111", salt)
	if err != nil {
		fmt.Printf("hash password failed, %v", err)
	}
	fmt.Printf("salt: %s, hashedPassword: %s\n", salt, hashedPassword)
}
