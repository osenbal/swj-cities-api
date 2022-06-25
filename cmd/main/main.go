package main

import (
	"net/http"
	"os"

	"gitlab.com/osenbal_slowly/swj-cities-api.git/pkg/routes"
)

func main() {

	e := routes.CreateCitiesRouters()

	if value, ok := os.LookupEnv("PORT"); ok {
		println(value)
		http.ListenAndServe(":"+os.Getenv("PORT"), e)
	} else {
		e.Start(":8080")
	}

	// if err != nil {
	// 	fmt.Println("SERVER RUNNING ON LOCAL")
	// }

	//local
}
