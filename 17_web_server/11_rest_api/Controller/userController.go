package Controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	database "github.com/ankitkumbhar/training/17_web_server/11_rest_api/Database"
)

type UsersStruct struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userId int
var name string
var email string

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Displaying from controller method")
	users := []UsersStruct{}
	db := database.Connect()
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := UsersStruct{}
	for rows.Next() {
		rows.Scan(&userId, &name, &email)
		user.Id = userId
		user.Name = name
		user.Email = email
		users = append(users, user)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	// var id string
	id := mux.Vars(r)["id"]
	users := []UsersStruct{}

	db := database.Connect()

	rows, err := db.Query("SELECT id, name, email from users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := UsersStruct{}
	if rows.Next() {
		rows.Scan(&userId, &name, &email)
		user.Id = userId
		user.Name = name
		user.Email = email

		users = append(users, user)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	user := UsersStruct{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)

	fmt.Println(user)
	db := database.Connect()
	stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user.Id = int(lastInsertedId)
	users := []UsersStruct{}
	users = append(users, user)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	db := database.Connect()

	_, err := db.Query("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	Index(w, r)
}

func Update(w http.ResponseWriter, r *http.Request) {
	user := UsersStruct{}
	id := mux.Vars(r)["id"]

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)

	db := database.Connect()
	_, err := db.Query("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		log.Fatal(err)
	}

	Index(w, r)
}
