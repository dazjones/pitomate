package dbtools

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/dazjones/pitomate/conf"
	"strings"
)

var db *sql.DB

func GetConnection() *sql.DB {

	db = *db

	if db != nil {
		return db
	}

	cnf := conf.GetConf()

	conStrings := []string{cnf.MySQLUser, ":", cnf.MySQLPass, "@tcp(",
		cnf.MySQLHost, ":", cnf.MySQLPort ,")/", cnf.MySQLName, "?charset=utf8"}

	db, _ = sql.Open("mysql", strings.Join(conStrings, ""))

	return db
}