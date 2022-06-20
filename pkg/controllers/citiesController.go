package controllers

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/repositories"
	. "gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/utils"
	"net/http"
)

func SearchByName(c echo.Context) error {
	cities, res := repositories.SearchByName(c)
	response := &Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    cities,
	}
	return c.JSON(http.StatusOK, response)
}
