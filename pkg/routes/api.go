package routes

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/controllers"
	"net/http"
)

// API Lists
var CreateCitiesRouters = func() *echo.Echo {
	route := echo.New()
	
	route.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Welcome to cities api (kelompok 9)")
	})

	route.GET("/api/city/search/:name", controllers.SearchByName)

	return route
}
