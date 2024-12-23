package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "go_db"
	port = 5432
	user = "postgres"
	password = "1234"
	dbname = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
							 host,   port,   host,   user,   password,   dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " +  dbname)

	return db, nil
}