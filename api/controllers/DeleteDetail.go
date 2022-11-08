package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Amulya301/todo-details/api/serializers"

	"github.com/Amulya301/todo-details/api/models"
	"github.com/Amulya301/todo-details/utils"
	"github.com/gorilla/mux"
)

func DeleteDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	detail := models.Details{}

	//get the slug by the parameter 'id'
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)
	todo, err1 := detail.RetrieveDetails(id)

	if err1 != nil {
		switch err1 {
		case utils.ErrResourceNotFound:
			utils.FindError(w, err1.Error(), http.StatusNotFound)
		default:
			utils.FindError(w, err1.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := detail.DeleteDetails(id)

	//if an error is found send it to the client and return
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.FindError(w, err.Error(), http.StatusNotFound)
		default:
			utils.FindError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	detailSerializer := serializers.DetailSerializer{
		Detail: []*models.Details{
			todo,
		},
		Many: false,
		StatusType: "OK",
		Code: 200,
	}

	resMap := map[string]interface{}{
		
		"code": detailSerializer.Code,
		"type": detailSerializer.StatusType,
		"delete": true,
		"data":  detailSerializer.Serialize()["data"],
	}

	// send the todo to the client
	_ = json.NewEncoder(w).Encode(resMap)

}
