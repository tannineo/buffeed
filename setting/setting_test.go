package setting_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/setting"
)

func TestConfig(t *testing.T) {
	Convey("Config read from config.json", t, func() {
		So(setting.Config.Port, ShouldEqual, 8080)
	})
}
