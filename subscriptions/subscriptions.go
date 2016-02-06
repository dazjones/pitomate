package subscriptions

import (
	dbtools "github.com/dazjones/pitomate/dbtools"
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

func Get(name string) *Subscription {
	cnf := conf.GetConf()

	db :=  dbtools.GetConnection()

	rows, err := db.Query("SELECT * FROM device_list WHERE name = ? AND hostname = ? LIMIT 1", name, cnf.Hostname)

	checkErr(err)

	s := new(Subscription)

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Name, &s.Gpio, &s.Pin, &s.Inverse, &s.Hostname)
	}

	return s
}

func GetAll() []Subscription {
	cnf := conf.GetConf()

	db :=  dbtools.GetConnection()

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