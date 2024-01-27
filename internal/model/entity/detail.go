package entity

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
)

type (
	Detail struct {
		*mysqlx.CommonDoc
		Name string `gorm:"column:name;type:varchar(2);default:'';comment:名字"`
	}
)
