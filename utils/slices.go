package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SliceSort[S ~[]E, E any](s S, less func(i, j int) bool) {
	sort.Slice(s, less)
}

func SliceContain[S ~[]E, E any](s S, contain func(e E) bool) bool {
	for _, v := range s {
		if contain(v) {
			return true
		}
	}
	return false
}

func SliceRemoveDuplicated[S ~[]E, E any](s S, compare func(e1 E, e2 E) bool) S {
	var ret S
	for _, e1 := range s {
		var find bool
		for _, e2 := range ret {
			if compare(e1, e2) {
				find = true
				break
			}
		}
		if find {
			continue
		}
		ret = append(ret, e1)
	}
	return ret
}

func SliceDistinct[S ~[]E, E any](s S, f func(e E) bool) (S, S) {
	var r1, r2 S
	for _, v := range s {
		if f(v) {
			t := v
			r1 = append(r1, t)
		} else {
			t := v
			r2 = append(r2, t)
		}
	}
	return r1, r2
}

func SliceMap[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func SliceExecutor[S ~[]E, RS ~[]R, E any, R any](s S, fn func(S) (RS, error), size int) (RS, error) {
	var ret RS
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		data, err := fn(s[i:end])
		if err != nil {
			return nil, err
		}
		ret = append(ret, data...)
	}
	return ret, nil
}

func MapKey2Slice[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// 字串轉 []uint64
func StringToUint64Slice(input string) ([]uint64, error) {
	if strings.TrimSpace(input) == "" {
		return []uint64{}, nil
	}

	strParts := strings.Split(input, ",")
	uintParts := make([]uint64, 0, len(strParts))

	for _, str := range strParts {
		num, err := strconv.ParseUint(strings.TrimSpace(str), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("can't transfer '%s' to uint64: %v", str, err)
		}
		uintParts = append(uintParts, num)
	}

	return uintParts, nil
}

// []uint64 轉字串
func Uint64ToStringSlice(slice []uint64) string {
	strParts := make([]string, len(slice))
	for i, num := range slice {
		strParts[i] = strconv.FormatUint(num, 10)
	}
	return strings.Join(strParts, ",")
}
