package rand_u

import (
	"math/rand"
	"bytes"
	"strings"
	"time"
)


const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 用掩码进行替换
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// 生成随机字符串
func RandStringWithLength(length int) string {
	b := make([]byte, length)

	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := length-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

var num = "0123456789"
var lower = "abcdefghijklmnopqrstuvwxyz"
var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// type is 0、a、A
// like StringWithType(64, "0aA")
func RandStringWithLengthAndType(length int, type_ string) (result string) {
	b := bytes.Buffer{}
	if strings.Contains(type_, "0") {
		b.WriteString(num)
	}
	if strings.Contains(type_, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(type_, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < length; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	return b.String()
}
