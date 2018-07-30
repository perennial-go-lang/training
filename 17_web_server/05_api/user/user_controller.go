package user

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/thedevsaddam/govalidator"
	"github.com/perennial-go-training/training/17_web_server/05_api/validator"
	"log"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintln(w, `{"error":"The user id must be a number"}`)
	} else {
		user := getUser(id)
		if user.Id == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, fmt.Sprintf(`{"error":"User with id %v not found"}`, id))
		} else {
			encoder := json.NewEncoder(w)
			encoder.Encode(&user)
		}
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	encoder := json.NewEncoder(w)
	encoder.Encode(&users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"fName": []string{"required", "alpha"},
		"lName": []string{"required", "alpha"},
		"age":   []string{"required", "numeric"},
	}

	messages := govalidator.MapData{
		"fName": []string{"required:First Name is required", "alpha:First Name can only contain characters"},
		"lName": []string{"required:Last Name is required", "alpha:Last Name can only contain characters"},
		"age":   []string{"required:Age is required", "number:Age must be a number"},
	}

	err := validator.Validate(rules, messages, r)
	if err != nil {
		log.Fatal(err)
	}

	user := &User{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)


}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Inside UpdateUser")
	fmt.Fprintln(w, "User Id", vars["id"])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Inside DeleteUser")
	fmt.Fprintln(w, "User Id", vars["id"])
}
