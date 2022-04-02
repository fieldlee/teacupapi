package utils

import (
	cryptorand "crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

const (
	letterBytes    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytesLen = len(letterBytes)
)

// GetRandNum 获取一个 min-max(不含max) 区间的随机数
func GetRandNum(min, max int) int {
	if max < min {
		tmp := max
		max = min
		min = tmp
	} else if min == max {
		return min
	}

	result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return min
	}

	num := int(result.Int64())

	return min + num
}

// 真随机
func RealRand(size int) string {
	var buf strings.Builder

	for i := 0; i < size; i++ {
		result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(letterBytesLen)))
		if err != nil {
			return ""
		}

		index := int(result.Int64())
		buf.WriteString(letterBytes[index : index+1])
	}
	str := buf.String()

	return str
}

// 真随机
func RealNumRand(size int) string {
	var buf strings.Builder

	for i := 0; i < size; i++ {
		result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(10))
		if err != nil {
			return ""
		}
		buf.WriteString(strconv.FormatInt(result.Int64(), 10))
	}
	str := buf.String()

	return str
}
