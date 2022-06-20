package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/models"
	"net/http"
)

func SearchByName(c echo.Context) (*[]models.City, *echo.HTTPError) {
	name := c.Param("name")
	city, db := models.SearchByName(name)

	if len(*city) == 0 || errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return city, echo.NewHTTPError(http.StatusNotFound)
	}

	return city, echo.NewHTTPError(http.StatusOK)
}
