package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

// Connect to database
func Connect() {
	connection, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=swj-cities-api sslmode=disable password=VINIX")
	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
