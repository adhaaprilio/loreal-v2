package config

import (
	"fmt"
	"loreal/entity"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port= %v sslmode=disable TimeZone=Asia/Shanghai", ENV.DB_HOST, ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_DATABASE, ENV.DB_PORT)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB = db
	log.Println("database connected")
	db.AutoMigrate(&entity.User{})
}

// func ConnectDB() {
// 	connStr := "user=postgres dbname=postgres"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Database Connected")
// 	db.Begin()
// }
