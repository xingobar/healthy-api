package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func init() {
	var err error

	Db, err = gorm.Open("mysql", "root:@/healthy?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
}
