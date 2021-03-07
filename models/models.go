package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

var Db *gorm.DB

func init() {
	var err error

	Db, err = gorm.Open("mysql", "root:@/healthy?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		fmt.Printf("page %d", page)

		if page == 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if limit == 0 {
			limit = 5
		}

		return db.Offset((page - 1) * limit).Limit(limit)
	}
}
