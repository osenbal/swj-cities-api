package main

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
)

func main() {
	r := echo.New()
	routes.CreateCitiesRouters(r)
	r.Start(":8080")
}
