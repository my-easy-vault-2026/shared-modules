package utils

import (
	"bytes"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"os"
	"shared-modules/common"
	"strconv"
	"testing"
	"time"
)

func Test_camelCase2underscoreCase(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{name: ""}, want: ""},
		{name: "2", args: args{name: "1"}, want: "1"},
		{name: "3", args: args{name: "a"}, want: "A"},
		{name: "4", args: args{name: "A"}, want: "A"},
		{name: "5", args: args{name: "abc"}, want: "ABC"},
		{name: "6", args: args{name: "123abc"}, want: "123ABC"},
		{name: "7", args: args{name: "abc123"}, want: "ABC123"},
		{name: "7", args: args{name: "abc123"}, want: "ABC123"},
		{name: "8", args: args{name: "abcDef"}, want: "ABC DEF"},
		{name: "9", args: args{name: "abCdEf"}, want: "AB CD EF"},
		{name: "10", args: args{name: "abc123Def"}, want: "ABC123 DEF"},
		{name: "11", args: args{name: "abc def"}, want: "ABC DEF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelCase2underscoreCase(tt.args.name); got != tt.want {
				t.Errorf("camelCase2underscoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateExcelWithFields(t *testing.T) {
	type row struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
		D string `json:"d"`
	}
	data := []row{
		{A: "A", B: "", C: "1", D: ""},
		{A: "B", B: "2", C: "3", D: ""},
	}
	filters := []string{"a", "c", "d"}
	columns := []string{"a1", "c3", "d4"}

	tmpF, err := os.CreateTemp("", "excel")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpF.Name())
	defer tmpF.Close()

	err = CreateExcelWithFields(data, filters, columns, tmpF)
	if err != nil {
		t.Fatal(err)
	}

	ef, err := excelize.OpenFile(tmpF.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer ef.Close()
	rows, err := ef.GetRows("Sheet1")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.Buffer{}
	for index, r := range rows {
		for _, c := range r {
			buf.WriteString(c)
		}
		t.Logf("index: %d, row: %s", index, buf.String())
		buf.Reset()
	}
}

func Test_deepSearchBaseTypes(t *testing.T) {
	t1 := 1
	t2 := "abc"
	t3 := true

	_, result := deepSearch(t1, "t1")
	t.Logf("result: %v", result)
	if result != "1" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	_, result = deepSearch(t2, "t2")
	t.Logf("result: %v", result)
	if result != "abc" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	_, result = deepSearch(t3, "t3")
	t.Logf("result: %v", result)
	if result != "true" {
		t.Errorf("deepSearch match fail: %v", result)
	}
}

func Test_deepSearchPointer(t *testing.T) {
	ts := "abc"

	_, result := deepSearch(&ts, "ts")
	t.Logf("result: %v", result)
	if result != "abc" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	var ts2 string

	_, result = deepSearch(&ts2, "ts2")
	t.Logf("result: %v", result)
	if result != "" {
		t.Errorf("deepSearch match fail: %v", result)
	}
}

func Test_deepSearchStruct(t *testing.T) {
	type obj struct {
		Z *bool   `json:"z"`
		A string  `json:"a"`
		B int     `json:"b"`
		C *string `json:"c"`
		D *int    `json:"d"`
	}

	tInt := 13
	tStr := "abc"

	to := obj{
		A: "ccc",
		B: 2,
		C: &tStr,
		D: &tInt,
	}

	ok, result := deepSearch(to, "a")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "ccc" {
		t.Errorf("deepSearch match fail: %v", result)
	}
	if result == "abc" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	ok, result = deepSearch(to, "b")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "2" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	ok, result = deepSearch(to, "c")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "abc" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	ok, result = deepSearch(to, "d")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "13" {
		t.Errorf("deepSearch match fail: %v", result)
	}
}

func Test_deepSearchGoDeeper(t *testing.T) {
	type o1 struct {
		X string `json:"x"`
		Y int    `json:"y"`
	}
	type o2 struct {
		G string `json:"o1"`
		H int    `json:"p1"`
	}
	type obj struct {
		A *o1 `json:"a"`
		B *o2 `json:"b"`
	}

	to := obj{
		A: &o1{
			X: "abc",
			Y: 111,
		},
		B: &o2{
			G: "def",
			H: 222,
		},
	}

	ok, result := deepSearch(to, "no")
	t.Logf("ok:%v, result:%v", ok, result)
	if ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}

	ok, result = deepSearch(to, "x")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "abc" {
		t.Errorf("deepSearch match fail: %v", result)
	}

	ok, result = deepSearch(to, "p1")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "222" {
		t.Errorf("deepSearch match fail: %v", result)
	}
}

func Test_deepSearchJumpCase(t *testing.T) {
	type o1 struct {
		X string `json:"x"`
		Y int    `json:"y"`
	}
	type o2 struct {
		G string `json:"o1"`
		H int    `json:"p1"`
	}
	type obj struct {
		A *o1 `json:"a"`
		B *o2 `json:"b"`
	}

	to := obj{
		A: &o1{
			X: "abc",
			Y: 111,
		},
		B: &o2{
			G: "def",
			H: 222,
		},
	}

	ok, result := deepSearch(to, "b")
	t.Logf("ok:%v, result:%v", ok, result)
	if !ok {
		t.Errorf("deepSearch find fail: %v", ok)
	}
	if result != "" {
		t.Errorf("deepSearch match fail: %v", result)
	}
}

func Test_CreateDynamicExcelWithFields(t *testing.T) {
	type o struct {
		A string  `json:"a"`
		B *string `json:"b"`
		C int     `json:"c"`
		D *int    `json:"d"`
		E *bool   `json:"e"`
	}

	tInt := 13
	tStr := "xyz"
	dChan := make(chan o)

	go func() {
		defer close(dChan)
		for i := 0; i < 1000; i++ {
			to := o{
				A: generateRandomString(10),
				B: &tStr,
				C: i,
				D: &tInt,
			}
			dChan <- to
		}
	}()

	filter := []string{"a", "b", "c", "d", "e"}
	columns := []string{"a", "b", "c", "d", "e"}

	f, err := os.Create("abc.xlsx")
	if err != nil {
		t.Fatalf("create excel failed: %v", err)
	}
	err = CreateDynamicExcelWithFields(dChan, filter, columns, f)
	if err != nil {
		t.Fatalf("create excel failed: %v", err)
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func Test_deepSearchCustomType(t *testing.T) {
	v := decimal.NewFromFloat(1.1)
	_, result := deepSearch(v, "")
	if result != "1.1" {
		t.Errorf("deepSearch fail: %v", result)
	}
}

func Test_deepSearch2String(t *testing.T) {
	result := cell2String(tcsx1)
	if result != "1" {
		t.Errorf("deepSearch fail: %v", result)
	}
	result = cell2String(tcsx2)
	if result != "1.23" {
		t.Errorf("deepSearch fail: %v", result)
	}
	result = cell2String(tcsx3)
	if result != "abc" {
		t.Errorf("deepSearch fail: %v", result)
	}
	result = cell2String(common.CURRENCY_USDT)
	if result != "usdt" {
		t.Errorf("deepSearch fail: %v", result)
	}
}

type customisedTestInt int
type customisedTestFloat float64
type customisedTestString string

const tcsx1 customisedTestInt = 1
const tcsx2 customisedTestFloat = 1.23
const tcsx3 customisedTestString = "abc"

func (c customisedTestInt) String() string {
	return strconv.Itoa(int(c))
}

func (c customisedTestFloat) String() string {
	return strconv.FormatFloat(float64(c), 'f', -1, 64)
}

func (c customisedTestString) String() string {
	return string(c)
}
