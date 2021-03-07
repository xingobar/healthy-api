package models

import (
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error

	Db, err = gorm.Open("mysql", "root:@/healthy")

	if err != nil {
		panic(err.Error())
	}
}
