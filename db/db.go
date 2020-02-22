package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func New() *gorm.DB {
	var err error

	DBMS     := "mysql"
	USER     := "root"
	PASS     := "####"
	PROTOCOL := "tcp(##.###.##.###:3306)"
	DBNAME   := "##"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	// CONNECT := ${DB_USER}:${DB_PASSWORD}@tcp(${DB_ADDRESS})/${DB_SCHEMA}?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true
	db,err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	return db
}
