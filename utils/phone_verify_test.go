package utils

import (
	"fmt"
	"testing"

	"github.com/nyaruka/phonenumbers"
)

func TestVerifyPhone(t *testing.T) {
	phone := "1534123456"
	regions := []string{"GB", "GG", "JE", "IM"}
	for _, region := range regions {
		num, err := phonenumbers.Parse(phone, region)
		fmt.Println("")
		fmt.Printf("%v: %v\n", num, err)
		if err != nil {
			fmt.Printf("號碼解析失敗: %v", err)
		}
		fmt.Println("")
		fmt.Printf("result: %t", phonenumbers.IsValidNumber(num))
	}

}
