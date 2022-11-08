package tests

import (
	"testing"

	"github.com/Amulya301/todo-details/cmd"
	"github.com/Amulya301/todo-details/utils"
)

func TestSeeder(t *testing.T) {
	t.Run("Seed todos", func(t *testing.T) {
		err := utils.SeedTodos(cmd.DbConnection)

		if err != nil {
			t.Fatal(err)
		}
	})
}
