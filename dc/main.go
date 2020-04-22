package main

import (
	"dataclips/Helpers"
	"dataclips/Routers"
	_ "dataclips/docs"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Dataclip post api
// @version 1.0
// @description This is a sample service for creating dataclips
// @termsOfService http://swagger.io/terms/
// @contact.name Awaniya Mohamed Elmabrouk
// @contact.email mohamed@craftfoundry.tech
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /
func main() {

	Helpers.Migration()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := Routers.InitRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
