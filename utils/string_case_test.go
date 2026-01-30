package utils

import (
	"context"
	"fmt"
	"testing"
)

func TestCamelToSnake(t *testing.T) {

	examples := []string{"SendOTPPassword", "sendOTPPassword", "SendOTP", "OTPSending"}
	for _, example := range examples {
		snakeCase := CamelToSnake(context.Background(), example)
		t.Log(fmt.Sprintf("%s -> %s\n", example, snakeCase))
	}
}
