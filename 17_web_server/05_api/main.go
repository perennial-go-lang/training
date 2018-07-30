package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/perennial-go-training/training/17_web_server/05_api/router"
	"github.com/perennial-go-training/training/17_web_server/05_api/user"
	"github.com/perennial-go-training/training/17_web_server/05_api/database"
)

const port = 6767

func main() {
	database.Start("root", "toor", "172.17.0.2", "3306", "resttest")
	defer database.Close()

	Router = mux.NewRouter()
	user.Bootstrap()

	log.Printf("Starting server at port %v", 6767)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), Router))
}