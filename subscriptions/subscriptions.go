package subscriptions

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
	"github.com/dazjones/pitomate/conf"
)

type Subscription struct {
	Id int
	Name string
	Gpio int
	Pin int
	Inverse bool
	Hostname string
}

func Get() []Subscription {
	cnf := conf.GetConf()

	conStrings := []string{cnf.MySQLUser, ":", cnf.MySQLPass, "@tcp(",
		cnf.MySQLHost, ":", cnf.MySQLPort ,")/", cnf.MySQLName, "?charset=utf8"}

	db, err := sql.Open("mysql", strings.Join(conStrings, ""))
	checkErr(err)

	rows, err := db.Query("SELECT * FROM device_list WHERE hostname = ?", cnf.Hostname)
	checkErr(err)

	var subs []Subscription

	for rows.Next() {
		s := new(Subscription)
		err = rows.Scan(&s.Id, &s.Name, &s.Gpio, &s.Pin, &s.Inverse, &s.Hostname)
		checkErr(err)
		subs = append(subs, *s)
	}

	return subs
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}