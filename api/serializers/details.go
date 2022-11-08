package serializers

import (
	"github.com/Amulya301/todo-details/api/models"
)

type DetailSerializer struct {
	Detail []*models.Details
	Many  bool
	Code int
	StatusType string
}

func (serializer *DetailSerializer) Serialize() map[string]interface{} {
	serializedData := make(map[string]interface{})

	DetailsArray := make([]interface{}, 0)
	for _, detail := range serializer.Detail {
		DetailsArray = append(DetailsArray, map[string]interface{}{
			"location":     	detail.Location,
			"description":   	detail.Description,
			"deadline" :		detail.Deadline,
		})
	}

	if serializer.Many {
		serializedData["data"] = DetailsArray
	} else {
		if len(DetailsArray) != 0 {
			serializedData["data"] = DetailsArray[0]
		} else {
			serializedData["data"] = make(map[string]interface{})
		}
	}

	return serializedData
}
