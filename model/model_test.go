package model_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
)

func TestMain(m *testing.M) {
	// before

	// run test
	m.Run()

	// after
	model.TearDownDB()
}

func Test_InitDB(t *testing.T) {
	Convey("InitDB...", t, func() {
		So(model.InitDB(), ShouldBeNil)

		// 删db
		Reset(func() {
			model.TearDownDB()
		})
	})
}
