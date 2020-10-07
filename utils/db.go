package utils

import (
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	//Db  *sql.DB
	Db  *gorm.DB
	err error
)

func init() {
	//Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/books")
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	// 打印出sql语句
	Db.LogMode(true)
	//
	//Db.SingularTable(true)

	// 连接池
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)


}
