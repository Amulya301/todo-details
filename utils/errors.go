package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ErrorResp struct {
	Type    string      `json:"type"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []interface{} `json:"errors,omitempty"`
}

var ErrResourceNotFound error = errors.New("status not found")

func FindError(w http.ResponseWriter, err interface{}, statusCode int) {
	errorResponse := ErrorResp{}
	errorResponse.Code = statusCode
	log.Println(err)
	if statusCode == http.StatusNotFound {
		errorResponse.Type = "Not Found"
		errorResponse.Message = "The specified resource was not found"
		errorResponse.Errors = append(errorResponse.Errors, err.(string))
	} else if statusCode == http.StatusBadRequest {
		errorResponse.Type = "Bad Request"
		errorResponse.Message = "One or more parameters are missing"
		errorResponse.Errors = append(errorResponse.Errors, err.(string))
	} else {
		errorResponse.Type = "Invalid request"
		errorResponse.Message = "Something went wrong"
		errorResponse.Errors = append(errorResponse.Errors, err.(string))
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(errorResponse)

}
