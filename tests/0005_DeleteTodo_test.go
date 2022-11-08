package tests

import (
	"net/http"
	"testing"

	"github.com/Amulya301/todo-details/utils"
)

func TestDeleteTodo(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad request with a non-existing todo", func(t *testing.T) {
		statusCode, _, _ := ts.Delete(t, "/todo/90", "")

		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})

	t.Run("Valid request", func(t *testing.T) {
		statusCode, _, _ := ts.Delete(t, "/todo/3", "")

		if statusCode != http.StatusNoContent {
			t.Errorf("want %d status code; got %d", http.StatusNoContent, statusCode)
		}
	})
}
