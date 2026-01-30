package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateDigit(f validator.FieldLevel) bool {
	val := f.Field().String()

	pattern := `^[0-9]+$`
	reg, err := regexp.Compile(pattern) // filter exclude chars
	if err != nil {
		return false
	}

	match := reg.MatchString(val)
	if !match {
		return false
	}
	return true
}

func ValidatePostal(f validator.FieldLevel) bool {
	val := f.Field().String()

	if len(val) > 10 {
		return false
	}

	pattern := `^[A-Za-z0-9\- ]+$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return reg.MatchString(val)
}
