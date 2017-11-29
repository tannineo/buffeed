package setting

import (
	"github.com/jinzhu/configor"
	"github.com/tannineo/buffeed/util"
)

// Config buffeed设置
var Config = struct {
	Port     uint   `default:"4000"`
	Salt     string `default:"233"`
	DataPath string
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
