package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Amulya301/todo-details/api/serializers"

	"github.com/Amulya301/todo-details/api/models"
	"github.com/Amulya301/todo-details/utils"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

	todo := models.Todo{}

	limit, offset := utils.ExtractPaginationParams(r)
	// get all todos
	todosList, err := todo.All( limit, offset)

	// if an error is found, sent the status to the client
	if err != nil {
		utils.FindError(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")

	todoSerializer := serializers.TodoSerializer{
		Todos: todosList,
		Many:  true,
		Code : http.StatusOK,
		StatusType : "OK",
	}

	resMap := map[string]interface{}{
		"code": todoSerializer.Code,
		"type": todoSerializer.StatusType,
		"data":  todoSerializer.Serialize()["data"],
	}

	// send the resMap to the client
	_ = json.NewEncoder(w).Encode(resMap)

}
