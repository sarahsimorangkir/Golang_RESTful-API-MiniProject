package main

import (
	"echo-rest/db"
	"echo-rest/openapi"
	"echo-rest/routes"
)

func main() {

	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))

	openapi.GetArt()

}
