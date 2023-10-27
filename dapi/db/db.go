package db

import (
	"dapi/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func Init() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "digimons.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&models.Digimon{})
	fmt.Println("Database Migrated")
}

func Close() {
	DBConn.Close()
}
