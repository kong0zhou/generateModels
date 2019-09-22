package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
)

var ConfViper *viper.Viper

func InitConf() (err error) {
	ConfViper = viper.New()
	ConfViper.SetConfigName(`conf`)
	ConfViper.AddConfigPath(`./`)
	ConfViper.SetConfigType(`yaml`)
	err = ConfViper.ReadInConfig()
	if err != nil {
		logs.Error(err)
		return
	}
	return nil
}
