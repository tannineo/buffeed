package model_test

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
)

func TestInitDB(t *testing.T) {
	Convey("InitDB...", t, func() {
		So(model.InitDB(), ShouldBeNil)

		// 删db
		Reset(func() {
			if _, err := os.Stat("data.db"); err == nil {
				os.Remove("data.db")
			}
		})
	})
}
