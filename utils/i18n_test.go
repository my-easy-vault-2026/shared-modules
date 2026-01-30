package utils

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
	"testing"
)

var langList = []string{"ar", "de", "en", "es", "fr", "ja", "ko", "pt", "ru", "vi", "zh_hans", "zh_hant"}

func Test18nQuery(t *testing.T) {
	err := InitI18N("../conf/")
	if err != nil {
		t.Error(err)
	}

	errCode := 10000190

	for _, lang := range langList {
		msg, err := queryErrMsg(lang, errCode)
		if err != nil {
			t.Errorf("lang: %s, err: %v", lang, err)
			continue
		}
		t.Logf("lang: %s, msg: %s", lang, msg)
	}
}

func queryErrMsg(lang string, code int) (string, error) {
	return i18n.NewLocalizer(Bundles[lang]).Localize(&i18n.LocalizeConfig{MessageID: "Code" + strconv.Itoa(code)})
}
