package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := setupRouter()
	r = postUser(r)
	r.Run(":8080")
}