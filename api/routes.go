package api

import (
	"github.com/Amulya301/todo-details/api/controllers"
	"github.com/gorilla/mux"
)

// Register all the routes for the application
var RouteTodos = func(router *mux.Router) *mux.Router {
	router.HandleFunc("/todo", controllers.GetAllTodos).Methods("GET").Name("todos.get-all")
	router.HandleFunc("/todo/{id}", controllers.GetTodoById).Methods("GET").Name("todos.get-single")
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST").Name("todos.add-single")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodo).Methods("PUT").Name("todos.update-single")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE").Name("todos.delete-single")

	router.HandleFunc("/detail", controllers.CreateDetail).Methods("POST").Name("details.add-single")
	router.HandleFunc("/detail/{id}", controllers.UpdateDetail).Methods("PUT").Name("details.update-single")
	router.HandleFunc("/detail/{id}", controllers.DeleteDetail).Methods("DELETE").Name("details.delete-single")


	return router
}
