package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

var Db *gorm.DB

// https://www.artacode.com/posts/sql/gorm-err/

func init() {
	var err error
	//"root:@/healthy?charset=utf8&parseTime=true"
	Db, err = gorm.Open("mysql","root:@/healthy?charset=utf8&parseTime=true")

	if err != nil {
		panic(err.Error())
	}
}

// 分頁
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		if page == 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if limit == 0 {
			limit = 20
		}

		return db.Offset((page - 1) * limit).Limit(limit)
	}
}
