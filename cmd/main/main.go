package main

import (
	"github.com/labstack/echo"
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
	"net/http"
	"os"
)

func main() {
	r := echo.New()
	routes.CreateCitiesRouters(r)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
