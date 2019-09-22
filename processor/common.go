package processor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/astaxie/beego/logs"
)

func TypeConvert(s string) (r string, err error) {
	if s == `` {
		err = fmt.Errorf(`s is nil`)
		logs.Error(err)
		return ``, err
	}
	if strings.Contains(s, `char`) || strings.Contains(s, `text`) {
		return `string`, nil
	}
	if strings.Contains(s, `int`) || strings.Contains(s, `serial`) {
		return `int`, nil
	}
	if strings.Contains(s, `numeric`) || strings.Contains(s, `decimal`) || strings.Contains(s, `real`) {
		return `float64`, nil
	}
	if in(s, []string{`bytea`}) {
		return `[]byte`, nil
	}
	if strings.Contains(s, `time`) || in(s, []string{`date`}) {
		return `time.Time`, nil
	}
	if strings.Contains(s, `bool`) {
		return `bool`, nil
	}
	return `interface{}`, nil
}

func in(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// 名字转化，将下划线转驼峰
func NameConvert(s string) (r string) {
	if s == `` {
		return ``
	}
	if s == `id` {
		return `ID`
	}
	ss := strings.Split(s, `_`)
	for _, v := range ss {
		if v == `` {
			continue
		}
		if v == `id` {
			r = r + `ID`
			continue
		}
		r = r + strings.Title(v)
	}
	return r
}

// 执行命令行
func execCommand(commandName string, params []string) (result string, err error) {
	if commandName == "" {
		err = fmt.Errorf("commandName is null")
		logs.Error(err)
		return "", err
	}
	cmd := exec.Command(commandName, params...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(err.Error() + `:` + stderr.String())
		return "", err
	}
	result = out.String()
	return
}
