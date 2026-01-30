package utils

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/maps"
)

// region 封裝式
type IHasExt interface {
	SetExt(map[string]interface{})
	GetExt() map[string]interface{}
	Res2Ext(context.Context, any)
}

type Ext struct {
	Ext map[string]interface{}
}

func (e *Ext) MarshalJSON() ([]byte, error) {

	if e.Ext == nil || len(e.Ext) == 0 {
		return []byte("{}"), nil
	}

	j, err := json.Marshal(e.Ext)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (e *Ext) MarshalText() (text []byte, err error) {
	return e.MarshalJSON()
}

func (e *Ext) SetExt(ext map[string]interface{}) {
	e.Ext = ext
}

func (e *Ext) GetExt() map[string]interface{} {
	return e.Ext
}

func (e *Ext) Res2Ext(ctx context.Context, t any) {
	ext := res2Ext(ctx, t)
	e.SetExt(ext)
}

func res2Ext(ctx context.Context, t any) map[string]interface{} {

	var tType = reflect.TypeOf(t)

	if t == reflect.Map {
		return nil
	}

	ext := make(map[string]interface{}, 0)
	var structType reflect.Type
	var structValue, structPtrValue reflect.Value
	if tType.Kind() == reflect.Struct {
		structType = reflect.TypeOf(t)
		structValue = reflect.ValueOf(t)
		structPtrValue = reflect.ValueOf(&t)
	} else if tType.Kind() == reflect.Pointer {
		structType = reflect.TypeOf(t).Elem()
		structValue = reflect.ValueOf(t).Elem()
		structPtrValue = reflect.ValueOf(t)
	}

	for i := 0; i < structType.NumField(); i++ {
		if structValue.Field(i).IsZero() {
			continue
		}
		// 下面這段是從json包複製過來的
		sf := structType.Field(i)
		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		for _, c := range name {
			switch {
			case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
				name = ""
				break
			case !unicode.IsLetter(c) && !unicode.IsDigit(c):
				name = ""
				break
			}
		}
		if name == "" {
			continue
		}
		ft := sf.Type
		if ft.Name() == "" && ft.Kind() == reflect.Pointer {
			// Follow pointer.
			ft = ft.Elem()
		}

		switch ft.Kind() {
		case reflect.Int64:
			ext[name] = strconv.FormatInt(structValue.Field(i).Int(), 10)
		case reflect.Uint64:
			ext[name] = strconv.FormatUint(structValue.Field(i).Uint(), 10)
		case reflect.Float64:
			ext[name] = strconv.FormatFloat(structValue.Field(i).Float(), 'f', 18, 64)
		case reflect.Pointer:
			switch structValue.Field(i).Elem().Kind() {
			case reflect.Int64:
				ext[name] = strconv.FormatInt(structValue.Field(i).Elem().Int(), 10)
			case reflect.Uint64:
				ext[name] = strconv.FormatUint(structValue.Field(i).Elem().Uint(), 10)
			case reflect.Float64:
				ext[name] = strconv.FormatFloat(structValue.Field(i).Elem().Float(), 'f', 18, 64)
			case reflect.Struct:
				ptr := structValue.Field(i).Interface()
				switch reflect.TypeOf(ptr) {
				case reflect.TypeOf((*decimal.Decimal)(nil)),
					reflect.TypeOf((*decimal.NullDecimal)(nil)):
					continue
				}
				child := res2Ext(ctx, ptr)
				if len(child) > 0 {
					ext[name] = res2Ext(ctx, ptr)
				}
			}
		case reflect.Struct:
			ptr := structPtrValue.Elem().Field(i).Addr().Interface()
			switch reflect.TypeOf(ptr) {
			case reflect.TypeOf((*decimal.Decimal)(nil)),
				reflect.TypeOf((*decimal.NullDecimal)(nil)):
				continue
			}
			child := res2Ext(ctx, ptr)
			if len(child) > 0 {
				ext[name] = res2Ext(ctx, ptr)
			}
		}
	}
	return ext
}

//endregion

// region 通用式
func MapWithExt(ctx context.Context, t any, pTag ...string) (map[string]interface{}, map[string]interface{}, error) {

	if t == nil {
		return nil, nil, nil
	}

	var tType = reflect.TypeOf(t)

	if t == reflect.Map {
		return map2Ext(ctx, t)
	}

	ori, ext := make(map[string]interface{}, 0), make(map[string]interface{}, 0)

	j, err := json.Marshal(t)
	if err != nil {
		return nil, nil, err
	}

	if err = json.Unmarshal(j, &ori); err != nil {
		return nil, nil, err
	}

	var structType reflect.Type
	var structValue, structPtrValue reflect.Value
	if tType.Kind() == reflect.Struct {
		structType = reflect.TypeOf(t)
		structValue = reflect.ValueOf(t)
		structPtrValue = reflect.ValueOf(&t)
	} else if tType.Kind() == reflect.Pointer {
		structType = reflect.TypeOf(t).Elem()
		structValue = reflect.ValueOf(t).Elem()
		structPtrValue = reflect.ValueOf(t)
	}

	for i := 0; i < structType.NumField(); i++ {
		if structValue.Field(i).IsZero() {
			continue
		}
		// 下面這段是從json包複製過來的
		stf := structType.Field(i)
		tag := stf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
	loop:
		for _, c := range name {
			switch {
			case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
				// name = ""
				// break loop
			case !unicode.IsLetter(c) && !unicode.IsDigit(c):
				name = ""
				break loop
			}
		}
		ft := stf.Type
		if ft.Name() == "" && ft.Kind() == reflect.Pointer {
			// Follow pointer.
			ft = ft.Elem()
		}

		isInherited := false
		if name == "" {
			if ft.Kind() != reflect.Struct {
				continue
			}
			name = pTag[0]
			isInherited = true
		}

		switch ft.Kind() {
		case reflect.Int64:
			ext[name] = strconv.FormatInt(structValue.Field(i).Int(), 10)
		case reflect.Uint64:
			ext[name] = strconv.FormatUint(structValue.Field(i).Uint(), 10)
		case reflect.Float64:
			ext[name] = strconv.FormatFloat(structValue.Field(i).Float(), 'f', 18, 64)
		case reflect.Pointer:
			switch structValue.Field(i).Elem().Kind() {
			case reflect.Int64:
				ext[name] = strconv.FormatInt(structValue.Field(i).Elem().Int(), 10)
			case reflect.Uint64:
				ext[name] = strconv.FormatUint(structValue.Field(i).Elem().Uint(), 10)
			case reflect.Float64:
				ext[name] = strconv.FormatFloat(structValue.Field(i).Elem().Float(), 'f', 18, 64)
			case reflect.Struct:
				ptr := structValue.Field(i).Interface()
				switch reflect.TypeOf(ptr) {
				case reflect.TypeOf((*decimal.Decimal)(nil)),
					reflect.TypeOf((*decimal.NullDecimal)(nil)):
					continue
				}
				cori, _, err := MapWithExt(ctx, ptr, name)
				if err != nil {
					return nil, nil, err
				}
				if len(cori) > 0 {
					ori[name] = cori
				}
			}
		case reflect.Struct:
			ptr := structPtrValue.Elem().Field(i).Addr().Interface()
			switch reflect.TypeOf(ptr) {
			case reflect.TypeOf((*decimal.Decimal)(nil)),
				reflect.TypeOf((*decimal.NullDecimal)(nil)),
				reflect.TypeOf((**decimal.Decimal)(nil)),
				reflect.TypeOf((**decimal.NullDecimal)(nil)):
				continue
			}
			cori, _, err := MapWithExt(ctx, ptr, name)
			if err != nil {
				return nil, nil, err
			}
			if len(cori) > 0 {
				if isInherited {
					maps.Copy(ori, cori)
				} else {
					ori[name] = cori
				}
			}
		case reflect.Slice:
			svf := structValue.Field(i)                        // 欄位指標
			st := ft.Elem()                                    // 欄位元素類型
			ss := make([]string, 0, svf.Len())                 // 新slice
			ms := make([]map[string]interface{}, 0, svf.Len()) // 新slice
			switch st.Kind() {
			case reflect.Int64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatInt(svf.Index(i).Int(), 10))
				}
				ext[name] = ss
			case reflect.Uint64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatUint(svf.Index(i).Uint(), 10))
				}
				ext[name] = ss
			case reflect.Float64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatFloat(svf.Index(i).Float(), 'f', 18, 64))
				}
				ext[name] = ss
			case reflect.Pointer:
				switch ft.Elem().Elem().Kind() {
				case reflect.Int64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatInt(svf.Index(i).Elem().Int(), 10))
					}
					ext[name] = ss
				case reflect.Uint64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatUint(svf.Index(i).Elem().Uint(), 10))
					}
					ext[name] = ss
				case reflect.Float64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatFloat(svf.Index(i).Elem().Float(), 'f', 18, 64))
					}
					ext[name] = ss
				case reflect.Struct:
					for i := 0; i < svf.Len(); i++ {
						ptr := svf.Index(i).Interface()
						switch reflect.TypeOf(ptr) {
						case reflect.TypeOf((*decimal.Decimal)(nil)),
							reflect.TypeOf((*decimal.NullDecimal)(nil)):
							continue
						}
						cori, _, err := MapWithExt(ctx, ptr, name)
						if err != nil {
							return nil, nil, err
						}
						if len(cori) > 0 {
							ms = append(ms, cori)
						}
					}
					ori[name] = ms

				}
			case reflect.Struct:
				ptr := structPtrValue.Elem().Field(i).Addr().Interface()
				switch reflect.TypeOf(ptr) {
				case reflect.TypeOf((*decimal.Decimal)(nil)),
					reflect.TypeOf((*decimal.NullDecimal)(nil)):
					continue
				}
				cori, _, err := MapWithExt(ctx, ptr, name)
				if err != nil {
					return nil, nil, err
				}
				if len(cori) > 0 {
					ori[name] = cori
				}
			}
		}
	}

	if len(ext) > 0 {
		ori["ext"] = ext
	}
	return ori, ext, nil
}

func map2Ext(ctx context.Context, m any) (map[string]interface{}, map[string]interface{}, error) {

	mv := reflect.ValueOf(m)

	if mv.Kind() != reflect.Map {
		return nil, nil, nil
	}

	if m == nil || mv.Len() == 0 {
		return nil, nil, nil
	}

	ori, ext := make(map[string]interface{}, 0), make(map[string]interface{}, 0)
	maps.Copy(m.(map[string]interface{}), ori)

	for _, k := range mv.MapKeys() {
		v := mv.MapIndex(k)
		switch t := v.Interface().(type) {
		case int64:
			ext[k.String()] = strconv.FormatInt(t, 10)
		case *int64:
			ext[k.String()] = strconv.FormatInt(*t, 10)
		case uint64:
			ext[k.String()] = strconv.FormatUint(t, 10)
		case *uint64:
			ext[k.String()] = strconv.FormatUint(*t, 10)
		case float64:
			ext[k.String()] = strconv.FormatFloat(t, 'f', 18, 64)
		case *float64:
			ext[k.String()] = strconv.FormatFloat(*t, 'f', 18, 64)
		case bool, *bool, []bool, []*bool,
			int, *int, []int, []*int,
			int8, *int8, []int8, []*int8,
			int16, *int16, []int16, []*int16,
			int32, *int32, []int32, []*int32,
			uint, *uint, []uint, []*uint,
			uint8, *uint8, []uint8, []*uint8,
			uint16, *uint16, []uint16, []*uint16,
			uint32, *uint32, []uint32, []*uint32,
			uintptr, *uintptr, []uintptr, []*uintptr,
			float32, *float32, []float32, []*float32,
			string, *string, []string, []*string,
			time.Time, *time.Time, []time.Time, []*time.Time,
			decimal.Decimal, *decimal.Decimal, []decimal.Decimal, []*decimal.Decimal,
			decimal.NullDecimal, *decimal.NullDecimal, []decimal.NullDecimal, []*decimal.NullDecimal:
		default:
			vt := v.Type()
			switch v.Kind() {
			case reflect.Pointer:
				cori, _, err := MapWithExt(ctx, v.Interface())
				if err != nil {
					return nil, nil, err
				}
				ori[k.String()] = cori
			case reflect.Struct:
				cori, _, err := MapWithExt(ctx, v.Addr().Interface())
				if err != nil {
					return nil, nil, err
				}
				ori[k.String()] = cori
			case reflect.Map:
				cori, _, err := map2Ext(ctx, v.Interface())
				if err != nil {
					return nil, nil, err
				}
				ori[k.String()] = cori
			case reflect.Slice, reflect.Array:
				s := make([]map[string]interface{}, v.Len())
				switch vt.Elem().Kind() {
				case reflect.Pointer:
					for i := 0; i < v.Len(); i++ {
						cori, _, err := MapWithExt(ctx, v.Index(i).Interface())
						if err != nil {
							return nil, nil, err
						}
						s = append(s, cori)
					}
					ori[k.String()] = s
				case reflect.Struct:
					for i := 0; i < v.Len(); i++ {
						cori, _, err := MapWithExt(ctx, v.Index(i).Addr().Interface())
						if err != nil {
							return nil, nil, err
						}
						s = append(s, cori)
					}
					ori[k.String()] = s
				case reflect.Map:
					for i := 0; i < v.Len(); i++ {
						cori, _, err := map2Ext(ctx, v.Index(i).Interface())
						if err != nil {
							return nil, nil, err
						}
						s = append(s, cori)
					}
					ori[k.String()] = s
				}
			default:
				return nil, nil, nil
			}
			cori, _, err := MapWithExt(ctx, t)
			if err != nil {
				return nil, nil, err
			}
			ori[k.String()] = cori
		}
	}

	if len(ext) > 0 {
		ori["ext"] = ext
	}
	return ori, ext, nil
}

func ToConverted(ctx context.Context, t any) (interface{}, error) {
	if t == nil {
		return nil, nil
	}

	var tType = reflect.TypeOf(t)

	switch tType.Kind() {
	case reflect.Int64:
		return strconv.FormatInt(t.(int64), 10), nil
	case reflect.Uint64:
		return strconv.FormatUint(t.(uint64), 10), nil
	case reflect.Float64:
		return strconv.FormatFloat(t.(float64), 'f', 18, 64), nil
	case reflect.Struct, reflect.Slice:
		p, s, m, err := ToConvertedMap(ctx, t)
		if err != nil {
			return nil, err
		}
		if s != nil {
			return s, nil
		}
		if m != nil {
			return m, nil
		}
		if p != nil {
			return p, nil
		}
	case reflect.Pointer:
		switch ti := t.(type) {
		case int64:
			return strconv.FormatInt(ti, 10), nil
		case *int64:
			return strconv.FormatInt(*ti, 10), nil
		case uint64:
			return strconv.FormatUint(ti, 10), nil
		case *uint64:
			return strconv.FormatUint(*ti, 10), nil
		case float64:
			return strconv.FormatFloat(ti, 'f', 18, 64), nil
		case *float64:
			return strconv.FormatFloat(*ti, 'f', 18, 64), nil
		case bool, *bool, []bool, []*bool,
			int, *int, []int, []*int,
			int8, *int8, []int8, []*int8,
			int16, *int16, []int16, []*int16,
			int32, *int32, []int32, []*int32,
			uint, *uint, []uint, []*uint,
			uint8, *uint8, []uint8, []*uint8,
			uint16, *uint16, []uint16, []*uint16,
			uint32, *uint32, []uint32, []*uint32,
			uintptr, *uintptr, []uintptr, []*uintptr,
			float32, *float32, []float32, []*float32,
			string, *string, []string, []*string,
			time.Time, *time.Time, []time.Time, []*time.Time,
			decimal.Decimal, *decimal.Decimal, []decimal.Decimal, []*decimal.Decimal,
			decimal.NullDecimal, *decimal.NullDecimal, []decimal.NullDecimal, []*decimal.NullDecimal:
			return t, nil
		default:
			p, s, m, err := ToConvertedMap(ctx, t)
			if err != nil {
				return nil, err
			}
			if s != nil {
				return s, nil
			}
			if m != nil {
				return m, nil
			}
			if p != nil {
				return p, nil
			}
		}
	}

	return t, nil
}

func ToConvertedMap(ctx context.Context, t any, pTag ...string) (interface{}, []interface{}, map[string]interface{}, error) {

	if t == nil {
		return nil, nil, nil, nil
	}

	var itf interface{}
	if itf = t; itf == nil {
		return nil, nil, nil, nil
	}

	var tType = reflect.TypeOf(t)

	if t == reflect.Map {
		m, err := map2ConvertedMap(ctx, t)
		return nil, nil, m, err
	}

	ori, ext := make(map[string]interface{}, 0), make(map[string]interface{}, 0) // slice出錯

	j, err := json.Marshal(t)
	if err != nil {
		return nil, nil, nil, err
	}
	if j[0] == byte('[') {
		// var structType reflect.Type
		var structValue, structPtrValue reflect.Value
		// structType = reflect.TypeOf(t)
		structValue = reflect.ValueOf(t)
		structPtrValue = reflect.ValueOf(&t)
		sori := make([]interface{}, 0, structValue.Len())
		for i := 0; i < structValue.Len(); i++ {
			element, err := ToConverted(ctx, structPtrValue.Elem().Elem().Index(i).Addr().Interface())
			if err != nil {
				return nil, nil, nil, err
			}
			sori = append(sori, element)
		}
		return nil, sori, nil, nil
	}

	if err = json.Unmarshal(j, &ori); err != nil {
		return nil, nil, nil, err
	}

	var structType reflect.Type
	var structValue, structPtrValue reflect.Value
	if tType.Kind() == reflect.Struct {
		structType = reflect.TypeOf(t)
		structValue = reflect.ValueOf(t)
		structPtrValue = reflect.ValueOf(&t)
	} else if tType.Kind() == reflect.Pointer {
		structType = reflect.TypeOf(t).Elem()
		structValue = reflect.ValueOf(t).Elem()
		structPtrValue = reflect.ValueOf(t)
	} else {
		structType = reflect.TypeOf(t)
		switch structType.Kind() {
		case reflect.Uint64:
			str := strconv.FormatUint(t.(uint64), 10)
			return str, nil, nil, nil
		case reflect.Float64:
			str := strconv.FormatFloat(t.(float64), 'f', 18, 64)
			return str, nil, nil, nil
		case reflect.Int64:
			str := strconv.FormatInt(t.(int64), 10)
			return str, nil, nil, nil
		}
		return nil, nil, nil, nil
	}

	for i := 0; i < structType.NumField(); i++ {
		if structValue.Field(i).IsZero() {
			continue
		}
		// 下面這段是從json包複製過來的
		stf := structType.Field(i)
		tag := stf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
	loop:
		for _, c := range name {
			switch {
			case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
				// name = ""
				// break loop
			case !unicode.IsLetter(c) && !unicode.IsDigit(c):
				name = ""
				break loop
			}
		}
		ft := stf.Type
		if ft.Name() == "" && ft.Kind() == reflect.Pointer {
			// Follow pointer.
			ft = ft.Elem()
		}

		isInherited := false
		if name == "" {
			if ft.Kind() != reflect.Struct {
				continue
			}
			name = pTag[0]
			isInherited = true
		}

		switch ft.Kind() {
		case reflect.Int64:
			ext[name] = strconv.FormatInt(structValue.Field(i).Int(), 10)
		case reflect.Uint64:
			ext[name] = strconv.FormatUint(structValue.Field(i).Uint(), 10)
		case reflect.Float64:
			ext[name] = strconv.FormatFloat(structValue.Field(i).Float(), 'f', 18, 64)
		case reflect.Pointer:
			switch structValue.Field(i).Elem().Kind() {
			case reflect.Int64:
				ext[name] = strconv.FormatInt(structValue.Field(i).Elem().Int(), 10)
			case reflect.Uint64:
				ext[name] = strconv.FormatUint(structValue.Field(i).Elem().Uint(), 10)
			case reflect.Float64:
				ext[name] = strconv.FormatFloat(structValue.Field(i).Elem().Float(), 'f', 18, 64)
			case reflect.Struct:
				ptr := structValue.Field(i).Interface()
				switch reflect.TypeOf(ptr) {
				case reflect.TypeOf((*decimal.Decimal)(nil)),
					reflect.TypeOf((*decimal.NullDecimal)(nil)):
					continue
				}
				p, s, m, err := ToConvertedMap(ctx, ptr, name)
				if err != nil {
					return nil, nil, nil, err
				}
				if m != nil {
					ori[name] = m
				}
				if s != nil {
					ori[name] = s
				}
				if p != nil {
					ori[name] = p
				}
			}
		case reflect.Struct:
			var ptr any
			if tType.Kind() == reflect.Struct {
				ptr = structPtrValue.Elem().Elem().Field(i).Interface()
			} else if tType.Kind() == reflect.Pointer {
				ptr = structPtrValue.Elem().Field(i).Interface()
			}
			switch reflect.TypeOf(ptr) {
			case reflect.TypeOf(decimal.Decimal{}),
				reflect.TypeOf(decimal.NullDecimal{}),
				reflect.TypeOf((*decimal.Decimal)(nil)),
				reflect.TypeOf((*decimal.NullDecimal)(nil)),
				reflect.TypeOf((**decimal.Decimal)(nil)),
				reflect.TypeOf((**decimal.NullDecimal)(nil)):
				continue
			}
			p, s, m, err := ToConvertedMap(ctx, ptr, name)
			if err != nil {
				return nil, nil, nil, err
			}
			if m != nil {
				if isInherited {
					maps.Copy(ori, m)
				} else {
					ori[name] = m
				}
			}
			if s != nil {
				ori[name] = s
			}
			if p != nil {
				ori[name] = p
			}
		case reflect.Slice:
			svf := structValue.Field(i)             // 欄位指標
			st := ft.Elem()                         // 欄位元素類型
			ss := make([]string, 0, svf.Len())      // 新slice
			ms := make([]interface{}, 0, svf.Len()) // 新slice
			switch st.Kind() {
			case reflect.Int64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatInt(svf.Index(i).Int(), 10))
				}
				ext[name] = ss
			case reflect.Uint64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatUint(svf.Index(i).Uint(), 10))
				}
				ext[name] = ss
			case reflect.Float64:
				for i := 0; i < svf.Len(); i++ {
					ss = append(ss, strconv.FormatFloat(svf.Index(i).Float(), 'f', 18, 64))
				}
				ext[name] = ss
			case reflect.Pointer:
				switch ft.Elem().Elem().Kind() {
				case reflect.Int64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatInt(svf.Index(i).Elem().Int(), 10))
					}
					ext[name] = ss
				case reflect.Uint64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatUint(svf.Index(i).Elem().Uint(), 10))
					}
					ext[name] = ss
				case reflect.Float64:
					for i := 0; i < svf.Len(); i++ {
						ss = append(ss, strconv.FormatFloat(svf.Index(i).Elem().Float(), 'f', 18, 64))
					}
					ext[name] = ss
				case reflect.Struct:
					for i := 0; i < svf.Len(); i++ {
						ptr := svf.Index(i).Interface()
						switch reflect.TypeOf(ptr) {
						case reflect.TypeOf((*decimal.Decimal)(nil)),
							reflect.TypeOf((*decimal.NullDecimal)(nil)):
							continue
						}
						p, s, m, err := ToConvertedMap(ctx, ptr, name)
						if err != nil {
							return nil, nil, nil, err
						}
						if m != nil {
							ms = append(ms, m)
						}
						if s != nil {
							ms = append(ms, s)
						}
						if p != nil {
							ms = append(ms, p)
						}
					}
					ori[name] = ms

				}
			case reflect.Struct:
				ptr := structPtrValue.Elem().Elem().Field(i).Interface()
				switch reflect.TypeOf(ptr) {
				case reflect.TypeOf((*decimal.Decimal)(nil)),
					reflect.TypeOf((*decimal.NullDecimal)(nil)):
					continue
				}
				p, s, m, err := ToConvertedMap(ctx, ptr, name)
				if err != nil {
					return nil, nil, nil, err
				}
				if m != nil {
					ori[name] = m
				}
				if s != nil {
					ori[name] = s
				}
				if p != nil {
					ori[name] = p
				}
			}
		}
	}

	for k, v := range ext {
		ori[k] = v
	}

	return nil, nil, ori, nil
}

func map2ConvertedMap(ctx context.Context, m any) (map[string]interface{}, error) {

	mv := reflect.ValueOf(m)

	if mv.Kind() != reflect.Map {
		return nil, nil
	}

	if m == nil || mv.Len() == 0 {
		return nil, nil
	}

	ori, ext := make(map[string]interface{}, 0), make(map[string]interface{}, 0)
	maps.Copy(m.(map[string]interface{}), ori)

	for _, k := range mv.MapKeys() {
		v := mv.MapIndex(k)
		switch t := v.Interface().(type) {
		case int64:
			ext[k.String()] = strconv.FormatInt(t, 10)
		case *int64:
			ext[k.String()] = strconv.FormatInt(*t, 10)
		case uint64:
			ext[k.String()] = strconv.FormatUint(t, 10)
		case *uint64:
			ext[k.String()] = strconv.FormatUint(*t, 10)
		case float64:
			ext[k.String()] = strconv.FormatFloat(t, 'f', 18, 64)
		case *float64:
			ext[k.String()] = strconv.FormatFloat(*t, 'f', 18, 64)
		case bool, *bool, []bool, []*bool,
			int, *int, []int, []*int,
			int8, *int8, []int8, []*int8,
			int16, *int16, []int16, []*int16,
			int32, *int32, []int32, []*int32,
			uint, *uint, []uint, []*uint,
			uint8, *uint8, []uint8, []*uint8,
			uint16, *uint16, []uint16, []*uint16,
			uint32, *uint32, []uint32, []*uint32,
			uintptr, *uintptr, []uintptr, []*uintptr,
			float32, *float32, []float32, []*float32,
			string, *string, []string, []*string,
			time.Time, *time.Time, []time.Time, []*time.Time,
			decimal.Decimal, *decimal.Decimal, []decimal.Decimal, []*decimal.Decimal,
			decimal.NullDecimal, *decimal.NullDecimal, []decimal.NullDecimal, []*decimal.NullDecimal:
		default:
			vt := v.Type()
			switch v.Kind() {
			case reflect.Pointer:
				p, s, m, err := ToConvertedMap(ctx, v.Interface())
				if err != nil {
					return nil, err
				}
				if m != nil {
					ori[k.String()] = m
				}
				if s != nil {
					ori[k.String()] = s
				}
				if p != nil {
					ori[k.String()] = p
				}
			case reflect.Struct:
				pc, sc, mc, err := ToConvertedMap(ctx, v.Addr().Interface())
				if err != nil {
					return nil, err
				}
				if mc != nil {
					ori[k.String()] = mc
				}
				if sc != nil {
					ori[k.String()] = sc
				}
				if pc != nil {
					ori[k.String()] = pc
				}
			case reflect.Map:
				cori, err := map2ConvertedMap(ctx, v.Interface())
				if err != nil {
					return nil, err
				}
				ori[k.String()] = cori
			case reflect.Slice, reflect.Array:
				s := make([]interface{}, v.Len())
				switch vt.Elem().Kind() {
				case reflect.Pointer:
					for i := 0; i < v.Len(); i++ {
						pc, sc, mc, err := ToConvertedMap(ctx, v.Index(i).Interface())
						if err != nil {
							return nil, err
						}
						if mc != nil {
							s[i] = mc
						}
						if sc != nil {
							s[i] = sc
						}
						if pc != nil {
							s[i] = pc
						}
					}
					ori[k.String()] = s
				case reflect.Struct:
					for i := 0; i < v.Len(); i++ {
						pc, sc, mc, err := ToConvertedMap(ctx, v.Index(i).Addr().Interface())
						if err != nil {
							return nil, err
						}
						if mc != nil {
							s = append(s, mc)
						}
						if sc != nil {
							s = append(s, sc)
						}
						if pc != nil {
							s = append(s, pc)
						}
					}
					ori[k.String()] = s
				case reflect.Map:
					for i := 0; i < v.Len(); i++ {
						cori, err := map2ConvertedMap(ctx, v.Index(i).Interface())
						if err != nil {
							return nil, err
						}
						s = append(s, cori)
					}
					ori[k.String()] = s
				}
			default:
				return nil, nil
			}
			pc, sc, mc, err := ToConvertedMap(ctx, t)
			if err != nil {
				return nil, err
			}
			if mc != nil {
				ori[k.String()] = mc
			}
			if sc != nil {
				ori[k.String()] = sc
			}
			if pc != nil {
				ori[k.String()] = pc
			}
		}
	}
	for k, v := range ext {
		ori[k] = v
	}

	return ori, nil
}

//endregion
