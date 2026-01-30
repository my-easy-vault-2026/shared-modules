package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"

	"github.com/xuri/excelize/v2"
)

// getExcelColumnName 將列索引轉換為 Excel 列名
func getExcelColumnName(index int) string {
	columnName := ""
	for index >= 0 {
		remainder := index % 26
		columnName = fmt.Sprintf("%c", 'A'+remainder) + columnName
		index = index/26 - 1
	}
	return columnName
}

func GetExcelColumnName(index int) string {
	columnName := ""
	for index >= 0 {
		remainder := index % 26
		columnName = fmt.Sprintf("%c", 'A'+remainder) + columnName
		index = index/26 - 1
	}
	return columnName
}

// CreateExcelBytesFromStruct 動態生成 Excel 文件並返回字節緩衝區
func CreateExcelBytesFromStruct(ctx context.Context, data interface{}) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// 反射獲取 struct 字段名稱作為表頭
	slice := reflect.ValueOf(data)
	if slice.Kind() != reflect.Slice || slice.Len() == 0 {
		return nil, fmt.Errorf("data must be a non-empty slice")
	}

	structValue := slice.Index(0)
	// 支持處理指標類型
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}
	if structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("slice elements must be structs or pointers to structs")
	}

	// 設定表頭
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structValue.Type().Field(i).Name
		col := getExcelColumnName(i) // 使用動態函式生成列名
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), fieldName)
	}

	// 填入數據
	for i := 0; i < slice.Len(); i++ {
		row := i + 2 // 從第 2 列開始填寫資料
		record := slice.Index(i)
		if record.Kind() == reflect.Ptr {
			record = record.Elem()
		}
		for j := 0; j < record.NumField(); j++ {
			col := getExcelColumnName(j) // 使用動態函式生成列名
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, row), record.Field(j).Interface())
		}
	}

	// 寫入到緩衝區
	buffer := new(bytes.Buffer)
	if err := f.Write(buffer); err != nil {
		return nil, fmt.Errorf("failed to write Excel to buffer: %w", err)
	}

	return buffer, nil
}

func CreateExcelWithFields[S ~[]E, E any](data S, filters []string, columns []string, writer io.Writer) error {
	if len(filters) != len(columns) {
		return fmt.Errorf("filters and columns must have the same length")
	}
	f := excelize.NewFile()
	defer f.Close()
	sheetName := "Sheet1"
	slice := reflect.ValueOf(data)
	if slice.Kind() != reflect.Slice || slice.Len() == 0 {
		return fmt.Errorf("data must be a non-empty slice")
	}
	structValue := slice.Index(0)
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs or pointers to structs")
	}
	for i := 0; i < len(columns); i++ {
		col := getExcelColumnName(i)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), camelCase2underscoreCase(columns[i]))
	}
	for i := 0; i < slice.Len(); i++ {
		row := i + 2
		record := slice.Index(i)
		if record.Kind() == reflect.Ptr {
			record = record.Elem()
		}
		for j := 0; j < len(filters); j++ {
			ok, cellVal := deepSearch(record.Interface(), filters[j])
			if ok {
				col := getExcelColumnName(j)
				_ = f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, row), fmt.Sprintf("%v", cellVal))
			}
		}
	}
	return f.Write(writer)
}

func CreateDynamicExcelWithFields[E any](elem <-chan E, filters []string, columns []string, writer io.Writer) error {
	if len(filters) != len(columns) {
		return fmt.Errorf("filters and columns must have the same length")
	}
	f := excelize.NewFile()
	defer f.Close()
	sheetName := "Sheet1"
	for i := 0; i < len(columns); i++ {
		col := getExcelColumnName(i)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), camelCase2underscoreCase(columns[i]))
	}
	var index int
	for e := range elem {
		te := reflect.ValueOf(e)
		row := index + 2
		if te.Kind() == reflect.Ptr {
			te = te.Elem()
		}
		for j := 0; j < len(filters); j++ {
			ok, cellVal := deepSearch(te.Interface(), filters[j])
			if ok {
				col := getExcelColumnName(j)
				_ = f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, row), fmt.Sprintf("%v", cellVal))
			}
		}
		index++
	}
	return f.Write(writer)
}

func deepSearch[E any](e E, target string) (bool, string) {
	if reflect.TypeOf(e).Kind() >= 1 && reflect.TypeOf(e).Kind() <= 14 || reflect.TypeOf(e).Kind() == reflect.String {
		objVal := reflect.ValueOf(e)
		fVal := objVal.Interface()
		return false, cell2String(fVal)
	} else if reflect.TypeOf(e).Kind() == reflect.Ptr {
		if reflect.ValueOf(e).IsNil() {
			return false, ""
		}
		e2 := reflect.ValueOf(e).Elem().Interface()
		return deepSearch(e2, target)
	} else if reflect.TypeOf(e).Kind() == reflect.Struct {
		if reflect.TypeOf(e).Name() == reflect.TypeOf(decimal.Decimal{}).Name() {
			return false, fmt.Sprintf("%v", reflect.ValueOf(e).Interface())
		} else if reflect.TypeOf(e).Name() == reflect.TypeOf(time.Time{}).Name() {
			tTime := reflect.ValueOf(e).Interface().(time.Time)
			return false, tTime.Format(time.RFC3339)
		} else {
			//fmt.Printf("unexpected field type %v\n", reflect.TypeOf(e).Name())
		}
		objVal := reflect.ValueOf(e)
		for i := 0; i < objVal.NumField(); i++ {
			fVal := objVal.Field(i).Interface()
			match, result := deepSearch(fVal, target)
			fName := objVal.Type().Field(i).Name
			fTag := objVal.Type().Field(i).Tag.Get("json")
			var find bool
			if strings.ToUpper(target) == strings.ToUpper(fName) {
				find = true
			} else if len(fTag) > 0 {
				ts := strings.Split(fTag, ",")
				for _, tss := range ts {
					if strings.ToUpper(target) == strings.ToUpper(tss) {
						find = true
						break
					}
				}
			}
			if find || match {
				return true, result
			}
		}
	}
	return false, ""
}

func cell2String(value interface{}) string {
	v := reflect.ValueOf(value)
	stringMethod := v.MethodByName("String")
	if stringMethod.IsValid() && stringMethod.Kind() == reflect.Func {
		results := stringMethod.Call(nil)
		if len(results) > 0 && results[0].Kind() == reflect.String {
			return results[0].Interface().(string)
		}
	}
	return fmt.Sprintf("%v", value)
}

func camelCase2underscoreCase(name string) string {
	rs := []rune(name)
	result := bytes.Buffer{}
	for _, ch := range rs {
		if unicode.IsUpper(ch) {
			if result.Len() > 0 {
				result.WriteRune(' ')
			}
			result.WriteRune(unicode.ToUpper(ch))
		} else {
			result.WriteRune(unicode.ToUpper(ch))
		}
	}
	return result.String()
}
