package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/babulal107/training/Assignment/09_website/database"
	"github.com/gorilla/mux"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Active    int    `json:"active"`
	CreatedAt string `json:"created_at"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := database.Db.Query("Select * from user")
	if err != nil {
		log.Fatal(err)
	}
	users := []User{}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Active, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// users := []User{
	// 	User{1, "Babulal", "babulal@gmail.com", "+91859569856", "Babulal123"},
	// 	User{2, "Rohit", "rohit@gmail.com", "+91859569856", "Rohit123"},
	// }
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(&users)
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	user := User{}
	err := database.Db.QueryRow("Select * from user WHERE id=?", vars["id"]).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Active, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(&user)
}
func Create(w http.ResponseWriter, r *http.Request) {

	user := User{}
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Phone = r.FormValue("phone")
	user.Password = r.FormValue("password")

	// Prepare statement for inserting data
	stmtIns, err := database.Db.Prepare("INSERT INTO user (name,email,phone,password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(user.Name, user.Email, user.Phone, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.ID = id
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(&user)
}

/**

**/
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := User{}
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Phone = r.FormValue("phone")
	user.Password = r.FormValue("password")
	// update
	stmt, err := database.Db.Prepare("update user set name=?, email=?,phone=?, password=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Phone, user.Password, vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	res := map[string]string{"status": strconv.FormatInt(int64(http.StatusOK), 10),
		"message": "User updated successfully"}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(&res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//delete
	stmt, err := database.Db.Prepare("update user set is_active=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("1", vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	res := map[string]string{"status": strconv.FormatInt(int64(http.StatusOK), 10),
		"message": "User deleted successfully"}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(&res)
}
