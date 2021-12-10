package zregex

import (
	"regexp"
)

func commonValid(regular string, target string) bool {
	rgx := regexp.MustCompile(regular)
	return rgx.MatchString(target)
}

// 验证手机号
func ValidTelephone(tel string) bool {
	return commonValid("^1[3456789]\\d{9}$", tel)
}

// 验证QQ
func ValidQQ(qq string) bool {
	return commonValid("^\\d{5,}$", qq)
}

// 检测验证码 6 位数字
func ValidCode(code string) bool {
	return commonValid("^\\d{6,}$", code)
}
