package setting

import (
	"github.com/jinzhu/configor"
	"github.com/tannineo/buffeed/util"
)

// Config buffeed设置
var Config = struct {
	Port     uint   `default:"4000"` // Port 占用端口
	Salt     string `default:"233"`  // Salt 密码摘要时用
	DataPath string // DataPath db存放路径
	Interval int    `default:"10"` // Interval feed拉取间隔 单位是minute
}{}

func init() {
	home, err := util.GetUserHome()
	if err != nil {
		panic(err)
	}
	configor.Load(&Config, home+"/config.json")
	if Config.DataPath == "" {
		home, err = util.GetUserHome()
		if err == nil {
			Config.DataPath = home + "/data.db"
		}
	}
}
