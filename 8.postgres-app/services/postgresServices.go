package services

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDb() {
	dsn := os.Getenv("PROSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost user=root password=example dbname=testdb port=5432 sslmode=disable"
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println("postgres connected!", DB)
}

func init() {
	connectDb()
	fmt.Println(DB)
}
