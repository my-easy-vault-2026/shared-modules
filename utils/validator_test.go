package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidator(t *testing.T) {

	va := validator.New()

	err := va.Var("https://www.google.com", "url")

	t.Log(err)
}
