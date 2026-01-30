package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundles map[string]*i18n.Bundle

func InitI18N(configPath string) error {
	Bundles = make(map[string]*i18n.Bundle, 0)

	// TODO: i18n setting
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "en.toml")
	Bundles["en"] = bundle
	Bundles["en_us"] = bundle
	Bundles["en-us"] = bundle
	Bundles["en_uk"] = bundle
	Bundles["en-uk"] = bundle

	bundle = i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "zh_hans.toml")
	Bundles["zh_cn"] = bundle
	Bundles["zh-cn"] = bundle
	Bundles["zh"] = bundle
	Bundles["zh_hans"] = bundle
	Bundles["zh-hans"] = bundle

	bundle = i18n.NewBundle(language.TraditionalChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "zh_hant.toml")
	Bundles["zh_tw"] = bundle
	Bundles["zh-tw"] = bundle
	Bundles["zh_hant"] = bundle
	Bundles["zh-hant"] = bundle

	bundle = i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "ja.toml")
	Bundles["ja"] = bundle

	bundle = i18n.NewBundle(language.Korean)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "ko.toml")
	Bundles["ko"] = bundle

	bundle = i18n.NewBundle(language.French)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "fr.toml")
	Bundles["fr"] = bundle

	bundle = i18n.NewBundle(language.German)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "de.toml")
	Bundles["de"] = bundle

	bundle = i18n.NewBundle(language.Spanish)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "es.toml")
	Bundles["es"] = bundle

	bundle = i18n.NewBundle(language.Portuguese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "pt.toml")
	Bundles["pt"] = bundle

	bundle = i18n.NewBundle(language.Vietnamese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "vi.toml")
	Bundles["vi"] = bundle

	bundle = i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "ru.toml")
	Bundles["ru"] = bundle

	bundle = i18n.NewBundle(language.Arabic)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(configPath + "ar.toml")
	Bundles["ar"] = bundle
	return nil
}

func Translate(ctx context.Context, msg common.TranslateMsg, value ...string) string {

	templateData := map[string]interface{}{}

	if len(value) > 0 {
		for i, v := range value {
			templateData[fmt.Sprintf("value%d", i+1)] = v
		}
	}

	bundle := Bundles["en"]
	c, ok := ctx.(*gin.Context)
	if ok {
		al := strings.ToLower(c.Request.Header.Get("Accept-Language"))
		if b, ok := Bundles[al]; ok {
			bundle = b
		}
	}

	str, err := i18n.NewLocalizer(
		bundle,
		language.English.String(),
	).Localize(&i18n.LocalizeConfig{
		MessageID:    string(msg),
		TemplateData: templateData,
	})

	if err != nil {
		logger.Errorf("text msg not set: [%s], ", string(msg), err)
		return string(msg)
	}

	return str
}

func TranslateByCode(ctx context.Context, language common.Language, code int, value ...string) string {

	templateData := map[string]interface{}{}

	if len(value) > 0 {
		for i, v := range value {
			templateData[fmt.Sprintf("value%d", i+1)] = v
		}
	}

	bundle := Bundles["en"]

	al := strings.ToLower(language.String())
	if b, ok := Bundles[al]; ok {
		bundle = b
	}

	str, err := i18n.NewLocalizer(
		bundle,
	).Localize(&i18n.LocalizeConfig{
		MessageID:    "Code" + strconv.Itoa(code),
		TemplateData: templateData,
	})

	if err != nil {
		// logger.Errorf("text msg not set: [%s], ", string(code), err)
		// return string(code)
	}

	return str
}

func MatchLang(header string) ([]string, error) {
	var matcher = language.NewMatcher([]language.Tag{
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
		language.SimplifiedChinese,
		language.TraditionalChinese,
		language.Japanese,
		language.Korean,
		language.French,
		language.German,
		language.Spanish,
		language.Portuguese,
		language.Russian,
		language.Vietnamese,
		language.Arabic,
	})

	tags, _, err := language.ParseAcceptLanguage(header)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(tags))
	for _, tag := range tags {
		if _, _, c := matcher.Match(tag); c == language.Exact {
			res = append(res, tag.String())
		}
	}
	return res, nil
}
