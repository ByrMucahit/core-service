package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error is occured on .env file please check")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER_NAME")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	fmt.Println("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)

	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database", err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}

}
