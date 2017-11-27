package setting

import "github.com/jinzhu/configor"

// Config buffeed设置
var Config = struct {
	Port uint `default:"4000"`
}{}

func init() {
	configor.Load(&Config, "../config.json")
}
