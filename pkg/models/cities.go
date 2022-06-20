package models

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/database"
)

var dbConnection *gorm.DB

type City struct {
	gorm.Model
	Name       string `gorm:"" json:"name" xml:"name" form:"name" query:"name"`
	Display    string `json:"display" xml:"display" form:"display" query:"display"`
	Latitude   string `json:"latitude" xml:"latitude" form:"latitude" query:"latitude"`
	Longtitude string `json:"longtitude" xml:"longtitude" form:"longtitude" query:"longtitude"`
}

func init() {
	database.Connect()
	dbConnection = database.GetDBConnection()
	dbConnection.AutoMigrate(&City{})
}

func SearchByName(name string) (*[]City, *gorm.DB) {
	var cities []City
	db := dbConnection.Where("name ILIKE ?", "%"+name+"%").Find(&cities)
	println(db.Value)
	return &cities, db
}
