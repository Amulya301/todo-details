package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Amulya301/todo-details/api/serializers"
	"github.com/Amulya301/todo-details/utils"

	"github.com/Amulya301/todo-details/api/models"
	"github.com/gorilla/mux"
)

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//write the header to the response
	w.Header().Set("Content-Type", "application/json")
	todo := models.Todo{}

	//get the slug by the parameter 'id'
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)
	

	// get the todo with this id first
	_, err := todo.Retrieve(id)

	if err != nil {
		utils.FindError(w, err, http.StatusNotFound)
	}
	
	json.NewDecoder(r.Body).Decode(&todo)

	todos, err := todo.Update(id)

	//if an error is found send it to the client and return
	if err != nil {
		utils.FindError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todoSerializer := serializers.TodoSerializer{
		Todos: []*models.Todo{
			todos,
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

	//Encode the created todos response to json and send it
	_ = json.NewEncoder(w).Encode(resMap)

}
