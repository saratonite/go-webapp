package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

type Todo struct {
	gorm.Model
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func InitDb() {
	connection, err := gorm.Open(sqlite.Open("./storage/data.db"), &gorm.Config{})

	Conn = connection

	if err != nil {
		log.Fatal("Error connecting database ")
	}

	fmt.Println("Database connected successfully")

	Conn.AutoMigrate(&Todo{})

	fmt.Println("Auto migrated")

}
