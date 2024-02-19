package entity

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
)

type (
	User struct {
		*mysqlx.CommonDoc
		Uuid     string `gorm:"column:uuid" db:"column:uuid" json:"uuid" form:"uuid"`
		Mobile   string `gorm:"column:mobile" db:"column:mobile" json:"mobile" form:"mobile"`
		Email    string `gorm:"column:email" db:"column:email" json:"email" form:"email"`
		State    int64  `gorm:"column:state" db:"column:state" json:"state" form:"state"`
		Password string `gorm:"column:password" db:"column:password" json:"password" form:"password"`
	}

	//DetailMongo struct {
	//	*mongox.CommonDoc `bson:",inline"`
	//	Name              string `bson:"name,omitempty"`
	//}
)
