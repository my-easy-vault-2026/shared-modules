package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

func GetRandString(n int) string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randStr := ""
	for i := 0; i < n; i++ {
		randIndex := r.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

func ToStr(value interface{}) string {
	var str string
	if value == nil {
		return str
	}
	switch v := value.(type) {
	case float64:
		str = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		str = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case int:
		str = strconv.Itoa(v)
	case uint:
		str = strconv.Itoa(int(v))
	case int8:
		str = strconv.Itoa(int(v))
	case uint8:
		str = strconv.Itoa(int(v))
	case int16:
		str = strconv.Itoa(int(v))
	case uint16:
		str = strconv.Itoa(int(v))
	case int32:
		str = strconv.Itoa(int(v))
	case uint32:
		str = strconv.Itoa(int(v))
	case int64:
		str = strconv.FormatInt(v, 10)
	case uint64:
		str = strconv.FormatUint(v, 10)
	case string:
		str = v
	case []byte:
		str = string(v)
	case time.Time:
		str = v.Format("2006-01-02 15:04:05")
	case decimal.Decimal:
		str = v.String()
	default:
		newValue, _ := json.Marshal(value)
		str = string(newValue)
	}
	return str
}

func Ptr[T any](t T) *T {
	return &t
}

func DecPtrToStr(dec *decimal.Decimal, places ...int) string {
	if dec != nil {
		if len(places) > 0 {
			return dec.StringFixed(int32(places[0]))
		}
		return dec.String()
	}
	return ""
}
