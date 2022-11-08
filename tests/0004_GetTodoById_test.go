package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Amulya301/todo-details/utils"
)

func TestGetTodoById(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad request with a non-existing todo", func(t *testing.T) {
		statusCode, _, _ := ts.Get(t, "/todo/90")

		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})

	t.Run("Valid request", func(t *testing.T) {
		statusCode, _, resBody := ts.Get(t, "/todo/1")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

			if response["name"] != "title1" {
				t.Errorf("want %s as the title; got %s", "title1", response["name"])
			}

			if response["completed"] != "Not done" {
				t.Errorf("want %s as the completed; got %s", "Not done", response["completed"])
			}
		}
	})
}
