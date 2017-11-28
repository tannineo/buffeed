package setting

import "github.com/jinzhu/configor"

// Config buffeed设置
var Config = struct {
	Port     uint   `default:"4000"`
	Salt     string `default:""`
	DataPath string `default:"/Users/carychan/Code/go/src/github.com/tannineo/buffeed/data.db"`
}{}

func init() {
	// TODO: 如何仅靠代码 区分测试正式环境? 运行测试时?
	configor.Load(&Config, "config.json")
}
