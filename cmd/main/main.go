package main

import (
	"net/http"
	"os"

	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
)

func main() {

	e := routes.CreateCitiesRouters()

	if _, ok := os.LookupEnv("PORT"); ok {
		// Deployment
		http.ListenAndServe(":"+os.Getenv("PORT"), e)
	} else {
		// Local
		e.Start(":8080")
	}

}
