package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/Amulya301/todo-details/cmd"
	"github.com/Amulya301/todo-details/utils"
)

//query := `
//CREATE TABLE IF NOT EXISTS todo (
//	id INT AUTO_INCREMENT,
//	title VARCHAR(100),
//	description VARCHAR(200),
//	completed VARCHAR(100),
//	PRIMARY KEY (id)
//);
//`

// Todo struct which describes the todo table in the database
type Todo struct {
	Id          int    `db:"id" json:"id"`
	Name 		string `db:"name" json:"name"`
	Completed   string `db:"completed" json:"completed"`
	Created_at   time.Time `db:"created_at" json:"created_at"`
	*Details
}

var TotalTodosCount int

// All - Retrieves all the records from the todo Table
// Params - todos ([]*Todo)
// Returns an error, if any

func (todo *Todo) All(limit int, offset int) ([]*Todo, error) {
	todosInstance := make([]*Todo, 0)
	/// execute the select query
	err := cmd.DbConnection.Select(&todosInstance, "SELECT * FROM todos t JOIN details d ON t.id = d.todoId LIMIT ? OFFSET ? ", limit, offset)

	// if an error is found, return it
	if err != nil {
		return nil, err
	}

	totalRecordsCountQuery := `SELECT COUNT(*) FROM todos`

	err = cmd.DbConnection.QueryRowx(totalRecordsCountQuery).Scan(&TotalTodosCount)
	if err != nil {
		return nil, err
	}

	return todosInstance, nil
}

// Retrieve - Retrieves a record from the Todo table
// Params - todo (*Todo), id int
// returns an error, if any
func (todo *Todo) Retrieve(id int) (*Todo, error) {
	//execute the query
	err := cmd.DbConnection.Get(todo, `SELECT * FROM todos t JOIN details d ON t.id = d.todoId WHERE t.id=? LIMIT 1`, id)

	//if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return todo, utils.ErrResourceNotFound
		default:
			return todo, err
		}
	}
	return todo, nil
}

// Insert - Inserts the value in the todo table
// params - todo (*Todo)
// Returns todo and error if any
func (todo *Todo) Insert() (*Todo, error) {

	log.Println("Creation Started..")

	tx := cmd.DbConnection.MustBegin()

	todo.Created_at = time.Now()
	todo.Completed = "To be done"
	// execute the insert query
	insertQuery := "insert into todos(name,created_at,completed) values (:name,:created_at,:completed);"
	insertQuery2 := "insert into details(todoId,location,description,deadline) values (:todoId,:location,:description,:deadline);"
	
	result, err := tx.NamedExec(insertQuery, &todo)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	todo.TodoId = int(lastId)

	_, err = tx.NamedExec(insertQuery2, &todo.Details)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// Commit the transaction to the database
	_ = tx.Commit()
	
	return todo, nil
}

// Update - Modifies the values of the specified record in the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (todo *Todo) Update(id int) (*Todo, error) {

	tx := cmd.DbConnection.MustBegin()
	updateQuery := `UPDATE todos SET name=:name , completed=:completed where id=:id`;
	//execute the query
	_, err := tx.NamedExec(updateQuery, todo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	updateQuery2 := `UPDATE details SET location=:location , deadline=:deadline, description=:description where id=:id`;
	//execute the query
	_, err = tx.NamedExec(updateQuery2, &todo.Details)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	updatedTodo := &Todo{}
	updatedTodoFromDB, err := updatedTodo.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return updatedTodoFromDB, nil
}

// Delete - Deletes the specified record from the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (todo *Todo) Delete(id int) (error) {

	deleteQuery := `DELETE FROM todos WHERE id=?`

	// execute the delete query
	result, err := cmd.DbConnection.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	numberOfRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// if numberOfRows affected is 0, return record not found error
	if numberOfRows == 0 {
		return utils.ErrResourceNotFound
	}

	return nil
}
