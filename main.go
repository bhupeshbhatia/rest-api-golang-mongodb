package main

import (
	"fmt"
	"net/http"

	db "database"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Starting the application...")
	fmt.Println(db.DbConn())

	router := mux.NewRouter()
	http.ListenAndServe(":8080", router)

}
