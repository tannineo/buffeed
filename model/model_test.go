package model_test

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
)

func TestMain(m *testing.M) {
	// before

	// run test
	result := m.Run()

	// after
	model.TearDownDB()

	// end
	os.Exit(result)
}

func Test_InitDB(t *testing.T) {
	Convey("InitDB...", t, func() {
		So(func() {
			model.NewDB()
		}, ShouldNotPanic)

		// 删db
		Reset(func() {
			model.TearDownDB()
		})
	})
}
