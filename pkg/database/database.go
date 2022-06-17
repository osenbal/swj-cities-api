package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	connection, err := gorm.Open("postgres", "host=ec2-34-225-159-178.compute-1.amazonaws.com port=5432 user=ywrabpnbrxmidb dbname=dde6pk2h71cob5 sslmode=disable password=1460639e3c459b49943a9447bb3e6fd0ff3e23a901f9d046fd3e4cfd895c5283")
	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
