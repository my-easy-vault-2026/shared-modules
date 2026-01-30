package utils

import (
	"shared-modules/logger"

	"github.com/nyaruka/phonenumbers"
)

func IsValidPhone(phone string, regions []string) bool {
	// 遍歷每個對應的地區進行驗證
	var flag = false
	for _, region := range regions {
		num, err := phonenumbers.Parse(phone, region)
		if err != nil {
			logger.Errorf("號碼解析失敗: %v", err)
			return false
		}
		if phonenumbers.IsValidNumber(num) {
			flag = true
			break
		}
	}

	// 檢查號碼是否有效
	return flag
}
