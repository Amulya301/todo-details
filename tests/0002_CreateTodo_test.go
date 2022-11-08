package tests

import (
	"net/http"
	"testing"

	"github.com/Amulya301/todo-details/utils"
)

func TestCreateTodo(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	// When the request body does not contain a description
	t.Run("Required Fields not Given", func(t *testing.T) {
		reqBody := `{
			"completed" : "In progress"
		}`

		statusCode, _, _ := ts.Post(t, "/todo", reqBody, "")

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	// When the request is valid with all fields
	t.Run("Valid request", func(t *testing.T) {
		reqBody := `{ 
				"name": "title 5",
				"completed" : "In progress"
			}`

		statusCode, _, _ := ts.Post(t, "/todo", reqBody, "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		}
	})
}
