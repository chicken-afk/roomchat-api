package main

import (
	"goboilerplate/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Call the Router function from router.go
	router.Router(r)
	r.Run(":8080")
}
