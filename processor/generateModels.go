package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"../common"
	"../models"
	"github.com/astaxie/beego/logs"
)

// 生成结构体需要的参数
type gStructArgs struct {
	Table      models.Table
	FolderName string
}

func GenerateModels() (err error) {
	tables, err := models.GetTables()
	if err != nil {
		logs.Error(err)
		return err
	}
	if tables == nil || len(tables) == 0 {
		err = fmt.Errorf(`tables is nil`)
		logs.Error(err)
		return err
	}
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error(err)
		return err
	}
	folderName := common.ConfViper.GetString(`folderName`)
	if folderName == `` {
		err = fmt.Errorf(`folderName is nil`)
		logs.Error(err)
		return err
	}
	err = os.MkdirAll(currentPath+`/outFiles/`+folderName, os.ModePerm)
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, v := range tables {
		file, err := os.Create(currentPath + `/outFiles/` + folderName + `/` + v.TableName + `.go`)
		if err != nil {
			logs.Error(err)
			return err
		}
		var SArg gStructArgs
		SArg.FolderName = folderName
		SArg.Table = v
		temp := template.New(`temp_struct.txt`)
		temp = temp.Funcs(template.FuncMap{`nameConvert`: NameConvert, `typeConvert`: TypeConvert})
		temp, err = temp.ParseFiles(currentPath + `/inFiles/temp_struct.txt`)
		if err != nil {
			logs.Error(err)
			return err
		}
		err = temp.Execute(file, SArg)
		if err != nil {
			logs.Error(err)
			return err
		}
		// go fmt filename.go 格式化文件
		result, err := execCommand(`go`, []string{`fmt`, currentPath + `/outFiles/` + folderName + `/` + v.TableName + `.go`})
		if err != nil {
			logs.Error(err)
			return err
		}
		logs.Info(`result:`, result)
	}
	return nil
}
