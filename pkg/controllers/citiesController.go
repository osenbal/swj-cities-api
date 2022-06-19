package controllers

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/repositories"
	. "gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/utils"
	"net/http"
)

func GetAllCity(c echo.Context) error {
	cities, res := repositories.GetAllCity()
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    *cities,
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func GetCityById(c echo.Context) error {
	city, res := repositories.GetCityById(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    city,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateCity(c echo.Context) error {
	city, res := repositories.InsertData(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    city,
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateCity(c echo.Context) error {
	city, res := repositories.UpdateData(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    city,
	}
	return c.JSON(http.StatusOK, response)
}

func SearchByName(c echo.Context) error {
	cities, res := repositories.SearchByName(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    cities,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteCity(c echo.Context) error {
	city, res := repositories.DeleteById(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    city,
	}
	return c.JSON(http.StatusOK, response)
}

func RestoreData(c echo.Context) error {
	city, res := repositories.RestoreData(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    city,
	}
	return c.JSON(http.StatusOK, response)
}

func PermanentDelete(c echo.Context) error {
	_, db := repositories.PermanentDelete(c)
	return c.JSON(http.StatusOK, db)
}
