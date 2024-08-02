package main

import (
	"github.com/pietroBragaAquinoJunior/go-gin-api/src/controllers"
)

func main() {
	r := controllers.SetupRouter()
	r = controllers.PostUser(r)
	r.Run(":8080")
}