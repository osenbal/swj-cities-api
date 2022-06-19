package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/config"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	var dsn string

	if config.Config["MODE"] == "production" {
		dsn = os.Getenv("DATABASE_URL")
	} else if config.Config["MODE"] == "local" {
		if config.Config["DB_DIALECT"] == "postgres" {
			dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=enable password=%s",
				config.Config["DB_HOST"], config.Config["DB_PORT"], config.Config["DB_USER"], config.Config["DB_NAME"], config.Config["DB_PASSWORD"])

		} else if config.Config["DB_DIALECT"] == "mysql" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.Config["DB_HOST"], config.Config["DB_PORT"], config.Config["DB_USER"], config.Config["DB_NAME"], config.Config["DB_PASSWORD"])
		}
	}

	connection, err := gorm.Open(config.Config["DB_DIALECT"], dsn)

	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDBConnection() *gorm.DB {
	return db
}
