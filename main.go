package main

import (
	"goboilerplate/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	/* Load ENV */
	errEnv := godotenv.Load()
	if errEnv != nil {
		logrus.Fatal(" LOAD ENV ", errEnv)
	}

	//gin mode
	if os.Getenv("SERVICE_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// Call the Router function from router.go
	router.Router(r)
	r.Run(":8080")
}
