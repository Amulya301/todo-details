package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Amulya301/todo-details/utils"
)

func TestUpdateTodo(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad request with a non-existing todo id", func(t *testing.T) {
		reqBody := `{
			"name" : "changed title",
			"completed": "Done"
		}`

		statusCode, _, _ := ts.Put(t, "/todo/90", reqBody)

		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})

	t.Run("Valid request", func(t *testing.T) {
		reqBody := `{
			"name": "changed title",
			"completed": "done"
		}`

		statusCode, _, resBody := ts.Put(t, "/todo/2", reqBody)

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

			if response["name"] != "changed title" {
				t.Errorf("want %s as title; got %s", "changed title", response["name"])
			}

			if response["completed"] != "done" {
				t.Errorf("want %s as completed; got %s", "done", response["completed"])
			}
		}
	})
}
