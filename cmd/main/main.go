package main

import (
	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
	"net/http"
	"os"
)

func main() {
	e := routes.CreateCitiesRouters()
	http.ListenAndServe(":"+os.Getenv("PORT"), e)

	//local
	//e.Start(":8080")
}
