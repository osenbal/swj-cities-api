package main

import (
	"net/http"
	"os"

	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
)

func main() {
	e := routes.CreateCitiesRouters()
	http.ListenAndServe(":"+os.Getenv("PORT"), e)

	//local
	//r.Start(":8080")
}
