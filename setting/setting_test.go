package setting_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/setting"
	"github.com/tannineo/buffeed/util"
)

// 没有db操作
// 不用tear down db
func Test_Config(t *testing.T) {
	Convey("Config read from config.json", t, func() {
		So(setting.Config.Port, ShouldEqual, 4000)
		So(setting.Config.Salt, ShouldEqual, "233")
		home, err := util.GetUserHome()
		if err != nil {
			panic(err)
		}
		So(setting.Config.DataPath, ShouldEqual, home+"/data.db")
	})
}
