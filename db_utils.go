package main

import (
// 	"github.com/jmoiron/sqlx"
	"flag"
	//"github.com/vharitonsky/iniflags"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db, preloadDB	*gorm.DB
	dbDriver = flag.String("DbDriver", "", "Database Driver")
	dbDSN = flag.String("DSN", "", "Database Data Source Name")
	
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
