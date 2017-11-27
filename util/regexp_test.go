package util_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/util"
)

var emailTestCases = []struct {
	testcase string
	is       string
	result   bool
}{
	{"fizzBuzz@gmail.com", " is", true},
	{"asdoqenovub-329@hotmail.com", " is", true},
	{"no#@O!UBD!DB(*BD#^@gg.cea", " is not", false},
	{"ohno-as.cxeq@foo.comeonbaby.heyhey", " is not", false},
}

func TestIsEmail(t *testing.T) {
	for _, tc := range emailTestCases {
		Convey(tc.testcase+tc.is+" an email address.", t, func() {
			So(util.IsEmail(tc.testcase), ShouldEqual, tc.result)
		})
	}
}

var pwdTestCases = []struct {
	testcase string
	is       string
	result   bool
}{
	{"fizzBuzz@gmail.com", " is not", false},
	{"foobar2000", " is", true},
	{"no#@O!UBD!Da", " is not", false},
	{"2333_wysiwyg", " is", true},
}

func TestIsPwd(t *testing.T) {
	for _, tc := range pwdTestCases {
		Convey(tc.testcase+tc.is+" a valid password.", t, func() {
			So(util.IsPwd(tc.testcase), ShouldEqual, tc.result)
		})
	}
}

var nicknameTestCases = []struct {
	testcase string
	is       string
	result   bool
}{
	{"fizzBuzz@gmail.com", " is not", false},
	{"foobar2000", " is", true},
	{"no#@O!UBD!Da", " is not", false},
	{"2333_wysiwyg", " is", true},
	{"2333_wysiwyg2333_wysiwyg", " is not", false},
	{"_", " is", true},
}

func TestIsNickname(t *testing.T) {
	for _, tc := range nicknameTestCases {
		Convey(tc.testcase+tc.is+" a valid nickname.", t, func() {
			So(util.IsNickname(tc.testcase), ShouldEqual, tc.result)
		})
	}
}
