package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

type Comlumn struct {
	ColumnNumber int
	ColumnName   string
	ColumnType   string
	NotNull      bool
	IsPrimaryKey bool
}

type Table struct {
	TableName string
	Comlumns  []Comlumn
}

func GetTables() (tables []Table, err error) {
	tables = make([]Table, 0)
	// 查询数据库中所有表
	tableNameRows, err := db.Query(findTablesSql, "r")
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	defer tableNameRows.Close()
	for tableNameRows.Next() {
		var table Table
		err = tableNameRows.Scan(&table.TableName)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		tables = append(tables, table)
	}
	if tables == nil || len(tables) == 0 {
		err = fmt.Errorf(`数据库中没有表`)
		logs.Error(err)
		return nil, err
	}
	// 查询数据库中所有视图
	viewNameRows, err := db.Query(findTablesSql, "v")
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	defer viewNameRows.Close()
	for viewNameRows.Next() {
		var table Table
		err = viewNameRows.Scan(&table.TableName)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		tables = append(tables, table)
	}

	for i := 0; i < len(tables); i++ {
		err := getColumn(&tables[i])
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	}
	return tables, nil
}

func getColumn(table *Table) (err error) {
	if table == nil {
		err = fmt.Errorf(`table is nil`)
		logs.Error(err)
		return err
	}
	if table.TableName == `` {
		err = fmt.Errorf(`table.TableName is nil`)
		logs.Error(err)
		return err
	}
	if table.Comlumns == nil {
		table.Comlumns = make([]Comlumn, 0)
	}
	comlumnRows, err := db.Query(findColumnsSql, table.TableName)
	if err != nil {
		logs.Error(err)
		return err
	}
	defer comlumnRows.Close()
	for comlumnRows.Next() {
		var comlumn Comlumn
		err = comlumnRows.Scan(&comlumn.ColumnNumber, &comlumn.ColumnName, &comlumn.NotNull, &comlumn.IsPrimaryKey, &comlumn.ColumnType)
		if err != nil {
			logs.Error(err)
			return err
		}
		table.Comlumns = append(table.Comlumns, comlumn)
	}
	return nil
}
