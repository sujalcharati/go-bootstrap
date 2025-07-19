package main

import (
	"fmt"
	"log"
	"net/http"

	"match-making/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/join", handlers.RegisterPlayer).Methods("POST")

	fmt.Println(" server is running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
