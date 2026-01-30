package utils

import (
	"context"
	"strings"
)

func CamelToSnake(ctx context.Context, str string) string {

	result := strings.Builder{}
	prevPrev := false
	prev := false
	for i, c := range str {
		if c >= 'A' && c <= 'Z' {
			if !prev && i != 0 {
				result.WriteRune('_')
			}
			result.WriteRune(c + 'a' - 'A')
			prevPrev = prev
			prev = true
		} else if prevPrev && prev {
			buf := result.String()
			lastRune := buf[len(buf)-1]
			buf = buf[:len(buf)-1]
			result.Reset()
			result.WriteString(buf)
			result.WriteRune('_')
			result.WriteByte(lastRune)
			result.WriteRune(c)
			prevPrev = prev
			prev = false
		} else {
			result.WriteRune(c)
			prevPrev = prev
			prev = false
		}
	}
	return result.String()
}
