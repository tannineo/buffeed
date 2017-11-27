package model_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
)

func TestInitDB(t *testing.T) {
	Convey("InitDB...", t, func() {
		So(model.InitDB(), ShouldBeNil)

		// 删db
		Reset(func() {
			model.TearDownDB()
		})
	})
}
