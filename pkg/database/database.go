package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/config"
)

var (
	db *gorm.DB
)

func Connect() {
	var dsn string

	if config.MODE == "production" {
		dsn = "postgres://ywrabpnbrxmidb:1460639e3c459b49943a9447bb3e6fd0ff3e23a901f9d046fd3e4cfd895c5283@ec2-34-225-159-178.compute-1.amazonaws.com:5432/dde6pk2h71cob5"
	} else if config.MODE == "local" {
		if config.DB_ENGINE == "postgres" {
			dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
				config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_TABLE, config.DB_PASSWORD)
		} else if config.DB_ENGINE == "mysql" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_TABLE, config.DB_PASSWORD)
		}
	}

	connection, err := gorm.Open(config.DB_ENGINE, dsn)

	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
