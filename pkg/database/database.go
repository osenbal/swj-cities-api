package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	connection, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
