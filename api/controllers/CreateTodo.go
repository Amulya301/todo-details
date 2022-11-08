package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Amulya301/todo-details/api/serializers"

	"github.com/Amulya301/todo-details/api/models"
	"github.com/Amulya301/todo-details/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// set header content type to application/json
	w.Header().Set("Content-Type", "application/json")
	todoInstance := models.Todo{}

	// decode the request body to todo
	json.NewDecoder(r.Body).Decode(&todoInstance)

	if todoInstance.Name == "" {
		utils.FindError(w, nil, http.StatusBadRequest)
		return
	}

	todo, err := todoInstance.Insert()
	
	// if an error is found, send it to the client and return
	if err != nil {
		if err == utils.ErrResourceNotFound{
			utils.FindError(w, err.Error(), http.StatusNotFound)
		} else{
			utils.FindError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Todo data serialization
	todoSerializer := serializers.TodoSerializer{
		Todos: []*models.Todo{
			todo,
		},
		Many: false,
		StatusType: "OK",
		Code: 200,
	}

	resMap := map[string]interface{}{
		"code": todoSerializer.Code,
		"type": todoSerializer.StatusType,
		"data":  todoSerializer.Serialize()["data"],
	}

	// send the created todo to the response
	_ = json.NewEncoder(w).Encode(resMap)
}
