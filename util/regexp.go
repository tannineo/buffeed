package util

import (
	"regexp"
)

var emailRegExp, pwdRegExp, nicknameRegExp *regexp.Regexp

func init() {
	// emailRegExp 定义了电子邮件的格式
	// See: https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/zh/01-chapter3.markdown#验证电子邮件地址
	emailRegExp, _ = regexp.Compile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)

	// pwdRegExp 定义了密码
	// 最小8位 最大20位
	// 只允许大小写 数字 下划线
	pwdRegExp, _ = regexp.Compile(`^(\w){8,20}$`)

	// nicknameRegExp 定义了昵称
	// 最小1位 最大20位
	// 只允许大小写 数字 下划线
	nicknameRegExp, _ = regexp.Compile(`^(\w){1,20}$`)
}

// IsEmail 验证是否是电子邮件
// 加上了字数限制
func IsEmail(s string) bool {
	if len([]rune(s)) > 128 {
		return false
	}
	return emailRegExp.MatchString(s)
}

// IsPwd 验证是否是密码
func IsPwd(s string) bool {
	return pwdRegExp.MatchString(s)
}

// IsNickname 验证是否是昵称
func IsNickname(s string) bool {
	return nicknameRegExp.MatchString(s)
}
