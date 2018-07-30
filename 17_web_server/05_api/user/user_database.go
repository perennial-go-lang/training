package user

import (
	"github.com/perennial-go-training/training/17_web_server/05_api/database"
	"log"
	"fmt"
)

func bootstrapDB() {
	_, err := database.DB.Query("CREATE TABLE IF NOT EXISTS users(id int not null auto_increment, first_name text not null," +
		"last_name text not null, age int not null, primary key(id))")
	if err != nil {
		log.Fatal(err)
	}

	var result int

	rows, err := database.DB.Query("SELECT count(*) FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result)
	}

	log.Print(result)

	if result == 0 {
		for i := 0; i < 100; i++ {
			_, err := database.DB.Query(fmt.Sprintf("INSERT INTO users (first_name, last_name, age)"+
				" VALUES('Person_%v', 'Name_%v', %v)", i, i, []int{24, 25, 26}[i%3]))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func getUser(id int) User {
	var (
		userId int
		fName string
		lName string
		age int
	)
	rows, err := database.DB.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %v", id))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userId, &fName, &lName, &age)
	}

	return User{userId, fName, lName, age}
}

func getUsers() Users {
	var (
		userId int
		fName string
		lName string
		age int
		users Users
	)

	rows, err := database.DB.Query(fmt.Sprintf("SELECT * FROM users"))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userId, &fName, &lName, &age)
		users = append(users, User{userId, fName, lName, age})
	}

	return users
}