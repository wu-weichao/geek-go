package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	NAME     = "root"
	PASSWORD = "123456"
	HOST     = "127.0.0.1"
	PORT     = "3307"
	DATABASE = "share"
	CHARSET  = "utf8"
)

var (
	db      *sql.DB
	dbError error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", NAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	db, dbError = sql.Open("mysql", dsn)
	if dbError != nil {
		log.Println("dsn: " + dsn)
		panic("数据库配置错误：" + dbError.Error())
	}
}
