package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"../common"
	"github.com/astaxie/beego/logs"
)

// 生成 init.go 需要的参数
type gInitArgs struct {
	FolderName   string
	PostgresArgs postgresArgs
}

type postgresArgs struct {
	Host     string
	Port     string
	User     string
	DBName   string
	SSLMode  bool
	Password string
}

func GenerateInit() (err error) {
	var args gInitArgs
	args.FolderName = common.ConfViper.GetString(`folderName`)
	args.PostgresArgs.Host = common.ConfViper.GetString(`postgresql.host`)
	args.PostgresArgs.Port = common.ConfViper.GetString(`postgresql.port`)
	args.PostgresArgs.User = common.ConfViper.GetString(`postgresql.user`)
	args.PostgresArgs.DBName = common.ConfViper.GetString(`postgresql.dbname`)
	args.PostgresArgs.SSLMode = common.ConfViper.GetBool(`postgresql.sslmode`)
	args.PostgresArgs.Password = common.ConfViper.GetString(`postgresql.password`)
	if args.FolderName == `` || args.PostgresArgs.Host == `` || args.PostgresArgs.Port == `` || args.PostgresArgs.User == `` || args.PostgresArgs.DBName == `` || args.PostgresArgs.Password == `` {
		err = fmt.Errorf(`配置文件中参数出现错误，生成参数有些为空`)
		logs.Error(err)
		return err
	}

	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error(err)
		return err
	}

	isExists, err := common.PathExists(currentPath + `/outFiles/` + args.FolderName)
	if err != nil {
		logs.Error(err)
		return err
	}

	if !isExists {
		err = os.MkdirAll(currentPath+`/outFiles/`+args.FolderName, os.ModePerm)
		if err != nil {
			logs.Error(err)
			return err
		}
	}
	file, err := os.Create(currentPath + `/outFiles/` + args.FolderName + `/init.go`)
	if err != nil {
		logs.Error(err)
		return err
	}

	temp, err := template.ParseFiles(currentPath + `/inFiles/temp_init.txt`)
	if err != nil {
		logs.Error(err)
		return err
	}
	err = temp.Execute(file, args)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
