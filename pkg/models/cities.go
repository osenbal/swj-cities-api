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

func (c *City) CreateCity() (*City, *gorm.DB) {
	dbConnection.NewRecord(c)
	db := dbConnection.Create(&c)
	return c, db
}

func SearchByName(name string) (*[]City, *gorm.DB) {
	var cities []City
	db := dbConnection.Where("name LIKE ?", "%"+name+"%").Find(&cities)
	return &cities, db
}

func GetAllCity() (*[]City, *gorm.DB) {
	var Cities []City
	db := dbConnection.Find(&Cities)
	return &Cities, db
}

func GetCityById(Id int64) (*City, *gorm.DB) {
	var getCity City
	db := dbConnection.Where("id=?", Id).Find(&getCity)
	return &getCity, db
}

func DeleteCity(Id int64) (*City, *gorm.DB) {
	var city City
	db := dbConnection.Where("id=?", Id).Find(&city)
	dbConnection.Where("id=?", Id).Delete(city)
	return &city, db
}

func RestoreData(Id int64) (*City, *gorm.DB) {
	city := new(City)
	db := dbConnection.Unscoped().Where("id=?", Id).Find(&city).Update("deleted_at", nil)
	return city, db
}

func PermanentDelete(Id int64) (*City, *gorm.DB) {
	var city City
	db := dbConnection.Unscoped().Where("id=?", Id).Find(&city)
	dbConnection.Unscoped().Where("id=?", Id).Delete(&city)
	return &city, db
}
