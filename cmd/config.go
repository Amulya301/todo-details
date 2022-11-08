package cmd

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DbConnection *sqlx.DB

func Connect() (*sqlx.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("No .env file found")
	}
	
	dbUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	//fmt.Println(dbUrl)
	dbConn, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err = dbConn.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT,
		name TEXT,
		completed TEXT,
		created_at DATE,
		PRIMARY KEY (id)
	);
	`
	_, err = dbConn.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	stmt := `
	CREATE TABLE IF NOT EXISTS details (
		id INT AUTO_INCREMENT,
		todoId INT AUTO_INCREMENT REFERENCES todos(id),
		location TEXT,
		description TEXT,
		deadline TEXT,
		PRIMARY KEY (id)
	);
	`
	_, err = dbConn.Exec(stmt)
	if err != nil {
		log.Fatalln(err)
	}

	return dbConn, nil
}
