package database

import (
	"fmt"
	"log"
	"os"

	"github.com/amandavmanduca/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connection := "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("DBPORT") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Error when connecting to Database")
	}
	fmt.Println("Connected to database")
	err := DB.AutoMigrate(&models.Student{})
	if err != nil {
		fmt.Printf("Error when migrating")
	}
}
