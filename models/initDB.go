package models

import (
	"database/sql"
	"fmt"

	"../common"
	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
)

var db *sql.DB

// 别忘记在main函数中 defer db.Close()

func InitDB() (err error) {
	host := common.ConfViper.GetString(`database.host`)
	port := common.ConfViper.GetString(`database.port`)
	user := common.ConfViper.GetString(`database.user`)
	dbname := common.ConfViper.GetString(`database.dbname`)
	sslmode := common.ConfViper.GetBool(`database.sslmode`)
	password := common.ConfViper.GetString(`database.password`)
	if host == `` || port == `` || user == `` || dbname == `` || password == `` {
		err = fmt.Errorf(`初始化数据库失败，参数为空`)
		logs.Error(err)
		return err
	}
	var dbInfo string
	if sslmode {
		dbInfo = fmt.Sprintf(`host=%s port=%s user=%s dbname=%s password=%s`, host, port, user, dbname, password)
	} else {
		dbInfo = fmt.Sprintf(`host=%s port=%s user=%s dbname=%s sslmode=disable password=%s`, host, port, user, dbname, password)
	}
	db, err = sql.Open(`postgres`, dbInfo)
	if err != nil {
		logs.Error(err)
		return
	}
	if err := db.Ping(); err != nil {
		logs.Error("连接数据库失败", err)
		return err
	}

	//SetMaxIdleConns设置空闲连接池中的最大连接数。
	db.SetMaxIdleConns(20)
	// 设置数据库的最大开放连接数。如果n≤0，则开放连接的数量没有限制。默认值为0
	db.SetMaxOpenConns(40)
	return nil
}
