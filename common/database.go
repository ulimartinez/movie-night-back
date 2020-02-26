package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "movieback:moviebackpass@/movie_nights?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
