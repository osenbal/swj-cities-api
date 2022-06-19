package routes

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/controllers"
	"net/http"
)

// API Lists
var CreateCitiesRouters = func() *echo.Echo {
	route := echo.New()
	route.GET("/api", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Welcome to cities api (kelompok 9)")
	})
	route.GET("/api/cities", controllers.GetAllCity)
	route.GET("/api/city/:id", controllers.GetCityById)
	route.GET("/api/city/search/:name", controllers.SearchByName)

	route.POST("/api/city", controllers.CreateCity)
	route.PUT("/api/city/:id", controllers.UpdateCity)
	route.PUT("/api/city/restore/:id", controllers.RestoreData)

	route.DELETE("/api/city/delete/:id", controllers.DeleteCity)
	route.DELETE("/api/city/permanent-delete/:id", controllers.PermanentDelete)
	return route
}
