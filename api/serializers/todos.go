package serializers

import (
	"github.com/Amulya301/todo-details/api/models"
)

type TodoSerializer struct {
	Todos []*models.Todo
	Many  bool
	Code int
	StatusType string
}

func (serializer *TodoSerializer) Serialize() map[string]interface{} {
	serializedData := make(map[string]interface{})

	todosArray := make([]interface{}, 0)
	for _, todo := range serializer.Todos {
		todosArray = append(todosArray, map[string]interface{}{
			"id":          todo.Id,
			"name":      	todo.Name,
			"completed":   	todo.Completed,
			"created_at":  todo.Created_at.Unix(),
			"details":    todo.Details,
		})
	}

	if serializer.Many {
		serializedData["data"] = todosArray
	} else {
		if len(todosArray) != 0 {
			serializedData["data"] = todosArray[0]
		} else {
			serializedData["data"] = make(map[string]interface{})
		}
	}

	return serializedData
}
