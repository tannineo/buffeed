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
		Convey(tc.testcase+tc.is+" a email address.", t, func() {
			So(util.IsEmail(tc.testcase), ShouldEqual, tc.result)
		})
	}
}
