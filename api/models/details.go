package models

import (
	"database/sql"
	"log"

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
type Details struct {
	Id              int    `db:"id" json:"id"`
	Location 		string `db:"location" json:"location"`
	Description     string `db:"description" json:"description"`
	Deadline   	    string `db:"deadline" json:"deadline"`
	TodoId			int  `db:"todoId" json:"todoId"`
}

var TotalDetailsCount int

// All - Retrieves all the records from the todo Table
// Params - todos ([]*Todo)
// Returns an error, if any

func (detail *Details) AllDetails(limit int, offset int) ([]*Details, error) {
	detailsInstance := make([]*Details, 0)
	/// execute the select query
	err := cmd.DbConnection.Select(&detailsInstance, "SELECT * FROM details LIMIT ? OFFSET ?", limit, offset)

	// if an error is found, return it
	if err != nil {
		return nil, err
	}

	totalRecordsCount := `SELECT COUNT(*) FROM details`

	err = cmd.DbConnection.QueryRowx(totalRecordsCount).Scan(&TotalDetailsCount)
	if err != nil {
		return nil, err
	}

	return detailsInstance, nil
}

// Retrieve - Retrieves a record from the Todo table
// Params - todo (*Todo), id int
// returns an error, if any
func (detail *Details) RetrieveDetails(id int) (*Details, error) {
	//execute the query
	err := cmd.DbConnection.Get(detail, `SELECT * FROM details WHERE id=? LIMIT 1`, id)

	//if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return detail, utils.ErrResourceNotFound
		default:
			return detail, err
		}
	}
	return detail, nil
}

// Insert - Inserts the value in the todo table
// params - todo (*Todo)
// Returns todo and error if any
func (detail *Details) InsertDetails() (*Details, error) {

	log.Println("Creation Started..")


	insertQuery := "insert into details(location,description,deadline,todoId) values (:location,:description,:deadline,:todoId);"
	// execute the insert query
	tx := cmd.DbConnection.MustBegin()
	_, err := tx.NamedExec(insertQuery, &detail)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	// Commit the transaction to the database
	_ = tx.Commit()


	return detail, nil
}

// Update - Modifies the values of the specified record in the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (detail *Details) UpdateDetails(id int) (*Details, error) {

	updateQuery := `UPDATE details SET location=:location , description=:description, deadline=:deadline where id=:id`;
	//execute the query
	tx := cmd.DbConnection.MustBegin()
	_, err := cmd.DbConnection.NamedExec(updateQuery, detail)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	updatedDetail := &Details{}
	updatedDetailFromDB, err := updatedDetail.RetrieveDetails(id)
	if err != nil {
		return nil, err
	}

	return updatedDetailFromDB, nil
}

// Delete - Deletes the specified record from the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (detail *Details) DeleteDetails(id int) (error) {

	deleteQuery := `DELETE FROM details WHERE id=?`

	// execute the delete query
	res, err := cmd.DbConnection.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	noOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// if numberOfRows affected is 0, return record not found error
	if noOfRows == 0 {
		return utils.ErrResourceNotFound
	}

	return nil
}
