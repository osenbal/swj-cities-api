package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/models"
	"net/http"
	"strconv"
)

func GetAllCity() (*[]models.City, *echo.HTTPError) {
	cities, db := models.GetAllCity()

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	return cities, echo.NewHTTPError(http.StatusOK)
}

func GetCityById(c echo.Context) (*models.City, *echo.HTTPError) {
	cityId := c.Param("id")
	ID, err := strconv.ParseInt(cityId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing city id")
	}

	city, db := models.GetCityById(ID)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	return city, echo.NewHTTPError(http.StatusOK)
}

func InsertData(c echo.Context) (*models.City, error) {
	newCity := &models.City{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&newCity)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	city, _ := newCity.CreateCity()
	return city, echo.NewHTTPError(http.StatusOK)
}

func UpdateData(c echo.Context) (*models.City, *echo.HTTPError) {
	defer c.Request().Body.Close()

	id := c.Param("id")
	ID, _ := strconv.ParseInt(id, 0, 0)

	city, db := models.GetCityById(ID)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	updateCity := &models.City{}
	err := json.NewDecoder(c.Request().Body).Decode(&updateCity)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if updateCity.Name != "" {
		city.Name = updateCity.Name
	}
	if updateCity.Display != "" {
		city.Display = updateCity.Display
	}
	if updateCity.Longtitude != "" {
		city.Longtitude = updateCity.Longtitude
	}
	if updateCity.Latitude != "" {
		city.Latitude = updateCity.Latitude
	}
	db.Save(&city)
	return city, echo.NewHTTPError(http.StatusOK)
}

func DeleteById(c echo.Context) (*models.City, *echo.HTTPError) {
	cityId := c.Param("id")
	ID, _ := strconv.ParseInt(cityId, 0, 0)

	city, db := models.GetCityById(ID)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	_, _ = models.DeleteCity(ID)

	return city, echo.NewHTTPError(http.StatusOK)
}

func SearchByName(c echo.Context) (*[]models.City, *echo.HTTPError) {
	name := c.Param("name")
	city, db := models.SearchByName(name)

	if len(*city) == 0 || errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return city, echo.NewHTTPError(http.StatusNotFound)
	}

	return city, echo.NewHTTPError(http.StatusOK)
}

func RestoreData(c echo.Context) (*models.City, *echo.HTTPError) {
	id := c.Param("id")
	ID, _ := strconv.ParseInt(id, 0, 0)

	city, db := models.RestoreData(ID)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	return city, echo.NewHTTPError(http.StatusOK)
}

func PermanentDelete(c echo.Context) (*models.City, *echo.HTTPError) {
	cityId := c.Param("id")
	ID, _ := strconv.ParseInt(cityId, 0, 0)

	delCity, db := models.PermanentDelete(ID)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}
	return delCity, echo.NewHTTPError(http.StatusOK)
}
