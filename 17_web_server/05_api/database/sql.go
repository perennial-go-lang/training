package database

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func getConnectionString(username, password, host, port, databaseName string) string {
return username + ":" + password +  "@tcp(" + host + ":" + port + ")/" + databaseName
}

func Start(username, password, host, port, databaseName string) {
	var err error
	DB, err = sql.Open("mysql", getConnectionString(username, password, host, port, databaseName))
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	DB.Close()
}

func Ping() bool {
	err := DB.Ping()
	if err != nil {
		defer log.Fatal(err)
		return false
	}
	return true
}