package tests

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Amulya301/todo-details/api"
	"github.com/Amulya301/todo-details/cmd"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Server *httptest.Server

var (
	Id          = ""
	Name 		= ""
	Completed   = ""
)

func TestMain(m *testing.M) {

	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file: ", err.Error())
		return
	}

	// Connect to test db
	cmd.DbConnection, err = cmd.Connect()
	dbConn := cmd.DbConnection

	if err != nil {
		log.Fatalln("Error connnecting to test db: ", err.Error())
		return
	}

	// Initialize new router for test
	router := mux.NewRouter()

	//// router.Route("/todos", api.RegisterTodoRoutes)
	api.RouteTodos(router)

	// Start new test server
	Server = httptest.NewServer(router)

	// Run tests
	exitVal := m.Run()

	// Close test server
	Server.Close()

	// utils.ClearTestDatabase(dbConn)
	// log.Print("Test Database cleared")

	// Close test database connection
	_ = dbConn.Close()

	// Exit main test
	os.Exit(exitVal)

}
