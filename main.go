package main

import (
	"log"
	"net/http"

	"github.com/Amulya301/todo-details/cmd"

	"github.com/Amulya301/todo-details/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("No .env file found")
	}

	dbConn, err := cmd.Connect()
	if err != nil {
		log.Fatalln("Error connecting db:", err.Error())
		return
	}
	cmd.DbConnection = dbConn

	api.RouteTodos(router)

	// http.Handle("/todos",)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", router))

}
