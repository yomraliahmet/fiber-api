package database

import (
	"log"
	"os"

	"github.com/yomraliahmet/fiber-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	var db *gorm.DB
	var err error

	switch os.Getenv("DB_CONNECTION") {
	case "mysql":
		dsn := os.Getenv("DB_DSN")
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	}

	//db, err = gorm.Open(sqlite.Open("fiberapi.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DbInstance{Db: db}

}

func AutoMigrate() {

	db := Database.Db

	log.Println("Running Migrations")
	// Add migrations
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
}
