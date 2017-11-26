package util

import (
	"regexp"
)

var emailRegExp *regexp.Regexp

func init() {
	emailRegExp, _ = regexp.Compile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
}

// IsEmail 验证是否是电子邮件
// See: https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/zh/01-chapter3.markdown#验证电子邮件地址
func IsEmail(s string) bool {
	return emailRegExp.MatchString(s)
}
