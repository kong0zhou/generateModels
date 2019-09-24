package main

import (
	"fmt"

	"./common"
	"./models"
	"./processor"
	"github.com/astaxie/beego/logs"
)

func main() {
	err := initEnv()
	if err != nil {
		logs.Error(err)
		return
	}
	err = processor.GenerateModels()
	if err != nil {
		logs.Error(err)
		return
	}
	err = processor.GenerateInit()
	if err != nil {
		logs.Error(err)
		return
	}
}

func initEnv() (err error) {
	err = common.InitLogger()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = common.InitConf()
	if err != nil {
		logs.Error(err)
		return err
	}
	err = models.InitDB()
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}
