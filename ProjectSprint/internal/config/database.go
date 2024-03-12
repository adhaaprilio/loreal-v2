package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", ENV.DB_HOST, ENV.DB_PORT, ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("database connected")

}
