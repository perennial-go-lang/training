package Employee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	DB "github.com/soniraju/training/Assignments/06_basic_api/DB"
)

type employee struct {
	Id    int    `json: "EmployeeId"`
	Name  string `json: "EmployeeName"`
	Email string `json: "EmployeeEmail"`
	City  string `json: "City"`
}

var EmpId int
var EmpName string
var EmpEmail string
var EmpCity string

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Index function..")
	employees := []employee{}
	db := DB.Connect()
	rows, err := db.Query("SELECT id, name, email, city FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer db.Close()

	employee := employee{}
	for rows.Next() {
		rows.Scan(&EmpId, &EmpName, &EmpEmail, &EmpCity)
		employee.Id = EmpId
		employee.Name = EmpName
		employee.Email = EmpEmail
		employee.City = EmpCity
		employees = append(employees, employee)
	}
	encoder := json.NewEncoder(w)
	// w.Header("content-type")
	err = encoder.Encode(employees)
	if err != nil {
		log.Fatalln(err)
	}

}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Show function..")

	id := mux.Vars(r)["id"]
	employees := []employee{}

	db := DB.Connect()
	rows, err := db.Query("SELECT id, name, email, city FROM employees WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer db.Close()

	employee := employee{}
	for rows.Next() {
		rows.Scan(&EmpId, &EmpName, &EmpEmail, &EmpCity)
		employee.Id = EmpId
		employee.Name = EmpName
		employee.Email = EmpEmail
		employee.City = EmpCity
		employees = append(employees, employee)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(employees)
	if err != nil {
		log.Fatalln(err)
	}

}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Store function..")
	emp := employee{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&emp)
	fmt.Println(emp)

	db := DB.Connect()
	stmt, err := db.Prepare("INSERT INTO employees (name,email,city) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	defer db.Close()

	result, err := stmt.Exec(emp.Name, emp.Email, emp.City)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	emp.Id = int(lastInsertedID)
	employees := []employee{}
	employees = append(employees, emp)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(employees)
	if err != nil {
		log.Fatal(err)
	}
}

func Destroy(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	fmt.Println("Inside Destroy function: ", id)
	db := DB.Connect()

	_, err := db.Query("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	Index(w, r)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("Inside Update function...")
	emp := employee{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&emp)

	db := DB.Connect()
	_, err := db.Query("UPDATE employees SET name = ?, email = ? , city = ? WHERE id = ?", emp.Name, emp.Email, emp.City, id)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Index(w, r)
}
