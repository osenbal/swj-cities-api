package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	connection, err := gorm.Open("postgres", "host=ec2-3-224-8-189.compute-1.amazonaws.com port=5432 user=nlbthlohzvexux dbname=d9afno0o6aoc9l sslmode=disable password=5c48088c7cfb36fe27a3b53d9acfa2e5a60c590e927c831ffde73fc85a94cb96")
	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
