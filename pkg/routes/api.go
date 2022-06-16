package routes

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/controllers"
)

var CreateCitiesRouters = func(route *echo.Echo) {
	route.GET("/api/cities", controllers.GetAllCity)
	route.GET("/api/city/:id", controllers.GetCityById)
	route.GET("/api/city/search/:name", controllers.SearchByName)

	route.POST("/api/city", controllers.CreateCity)
	route.PUT("/api/city/:id", controllers.UpdateCity)
	route.PUT("/api/city/restore/:id", controllers.RestoreData)

	route.DELETE("/api/city/delete/:id", controllers.DeleteCity)
	route.DELETE("/api/city/permanent-delete/:id", controllers.PermanentDelete)
}
