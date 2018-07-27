package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func getConnection(Username, Password, DatabaseName string) (con string) {

	return (Username + ":" + Password + "@/" + DatabaseName)
}

func Connect() {
	var err error
	Db, err = sql.Open("mysql", getConnection("root", "root", "go"))
	if err != nil {
		panic(err.Error())
	}
}
