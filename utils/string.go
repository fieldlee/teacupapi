package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 获取source的子串,如果start小于0或者end大于source长度则返回""
// start:开始index，从0开始，包括0
// end:结束index，以end结束，但不包括end
func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

// FloatPrecision float 精度转换
func FloatPrecision(fStr string, prec int, round bool) (string, error) {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return "", err
	}

	f = Precision(f, prec, round)
	str := strconv.FormatFloat(f, 'f', prec, 64)

	return str, nil
}

// FloatPrecisionStr float 转换为 string 精度转换
func FloatPrecisionStr(f float64, prec int, round bool) string {
	ff := Precision(f, prec, round)
	str := strconv.FormatFloat(ff, 'f', prec, 64)

	return str
}

// FloatPrecision float 精度转换
func FloatFPrecision(fStr string, prec int, round bool) (float64, error) {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return 0, err
	}

	return Precision(f, prec, round), nil
}

// Precision 支持精度以及是否四舍五入, round: true 为四舍五入, false 不是四舍五入
func Precision(f float64, prec int, round bool) float64 {
	// 需要加上对长度的校验, 否则直接用 math.Trunc 会有bug(1.14会变成1.13)
	arr := strings.Split(strconv.FormatFloat(f, 'f', -1, 64), ".")
	if len(arr) < 2 {
		return f
	}
	if len(arr[1]) <= prec {
		return f
	}
	pow10N := math.Pow10(prec)

	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}

	return math.Trunc((f)*pow10N) / pow10N
}

//生成随机字符串
func GetRandomString(length int) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func MapToQuery(m map[string]string) string {
	var sBuild strings.Builder
	for k, v := range m {
		if sBuild.Len() > 0 {
			sBuild.WriteString("&")
		}
		sBuild.WriteString(fmt.Sprint(k, "=", v))
	}
	return sBuild.String()
}

func Md5EncodeToString(s string) string {
	hexCode := md5.Sum([]byte(s))
	return hex.EncodeToString(hexCode[:])
}
