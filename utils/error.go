package utils

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func WrapBusinessError(ctx context.Context, err error) *common.BusinessError {
	if err != nil {
		err1, ok := err.(*common.BusinessError)
		if ok {
			return err1
		} else {
			return NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
		}
	}
	return nil
}

func GetErrorMsg(ctx context.Context, code int, lang string) (msg string) {
	if lang == "" {
		lang = strings.ToLower(fmt.Sprint(ctx.Value(common.HEADER_ACCEPT_LANGUAGED)))
	}
	bundle, ok := Bundles[lang]
	if !ok || lang == "" {
		bundle = Bundles["en"]
	}

	eMsg, err := i18n.NewLocalizer(
		bundle,
	).Localize(&i18n.LocalizeConfig{
		MessageID: "Code" + strconv.Itoa(code),
	})

	// 如果該語言找不到對應，則以英文為主
	if err != nil {
		logger.Warnf("text msg not set for code %d in bundle %v", code, err)
		bundle = Bundles["en"]
		eMsg, err = i18n.NewLocalizer(
			bundle,
		).Localize(&i18n.LocalizeConfig{
			MessageID: "Code" + strconv.Itoa(code),
		})
		if err != nil {
			logger.Warnf("text msg not set for code %d in bundle %v", code, err)
			eMsg = "error message not set"
		}
		eMsg = "error message not set"
	}

	return eMsg
}

func NewBusinessErrorWithArg(ctx context.Context, code int, arg map[common.BusinessErrorArg]string) (ret *common.BusinessError) {

	lang := strings.ToLower(fmt.Sprint(ctx.Value(common.HEADER_ACCEPT_LANGUAGED)))
	bundle, ok := Bundles[lang]
	if !ok || lang == "" {
		bundle = Bundles["en"]
	}

	eMsg, err := i18n.NewLocalizer(
		bundle,
	).Localize(&i18n.LocalizeConfig{
		MessageID: "Code" + strconv.Itoa(code),
	})

	// 如果該語言找不到對應，則以英文為主
	if err != nil {
		logger.Warnf("text msg not set for code %d in bundle %v", code, err)
		bundle = Bundles["en"]
		eMsg, err = i18n.NewLocalizer(
			bundle,
		).Localize(&i18n.LocalizeConfig{
			MessageID: "Code" + strconv.Itoa(code),
		})
		if err != nil {
			logger.Warnf("text msg not set for code %d in bundle %v", code, err)
			eMsg = "error message not set"
		}
		eMsg = "error message not set"
	}

	for k, v := range arg {
		eMsg = strings.ReplaceAll(eMsg, "{$"+string(k)+"}", v)
	}

	ret = common.NewBusinessError(
		code,
		eMsg,
	)
	return
}

func NewBusinessError(ctx context.Context, code int, msg ...string) (ret *common.BusinessError) {
	if len(msg) != 0 {
		return common.NewBusinessError(code, msg[0])
	}

	lang := strings.ToLower(fmt.Sprint(ctx.Value(common.HEADER_ACCEPT_LANGUAGED)))
	bundle, ok := Bundles[lang]
	if !ok || lang == "" {
		bundle = Bundles["en"]
	}

	eMsg, err := i18n.NewLocalizer(
		bundle,
	).Localize(&i18n.LocalizeConfig{
		MessageID: "Code" + strconv.Itoa(code),
	})

	// 如果該語言找不到對應，則以英文為主
	if err != nil {
		logger.Warnf("text msg not set for code %d in bundle %v", code, err)
		bundle = Bundles["en"]
		eMsg, err = i18n.NewLocalizer(
			bundle,
		).Localize(&i18n.LocalizeConfig{
			MessageID: "Code" + strconv.Itoa(code),
		})
		if err != nil {
			logger.Warnf("text msg not set for code %d in bundle %v", code, err)
			eMsg = "error message not set"
		}
		eMsg = "error message not set"
	}

	ret = common.NewBusinessError(
		code,
		eMsg,
	)
	return
}

func EqualBusinssError(code int, err error) bool {
	if bErr, ok := err.(*common.BusinessError); ok {
		return bErr.Code() == code
	}
	return false
}

func FormatErrorChainArrow(err error) string {
	return FormatErrorChain(err, " -> ")
}

func FormatErrorChainLink(err error) string {
	return FormatErrorChain(err, "🔗")
}

func FormatErrorChain(err error, separator string) string {
	var result strings.Builder

	for err != nil {
		result.WriteString(err.Error())
		err = errors.Unwrap(err)
		if err != nil {
			result.WriteString(separator)
		}
	}

	return result.String()
}

func FormatErrorChainHierarchical(err error) string {
	var result strings.Builder
	level := 0

	for err != nil {
		if level > 0 {
			result.WriteString("\n")
			result.WriteString(strings.Repeat("  ", level)) // 縮進
			result.WriteString("└─ ")
		}
		result.WriteString(err.Error())
		err = errors.Unwrap(err)
		level++
	}

	return result.String()
}

func GetRootCause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
