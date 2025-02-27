package main

import (
	"chatroom-api/database"
	"chatroom-api/entities"
	"chatroom-api/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	//Run gorm auto migrate
	database.SetupDatabaseConnection().AutoMigrate(
		&entities.User{},
		&entities.Roomchat{},
		&entities.RoomchatUser{},
	)

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

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-key")

		c.Next()
	})

	// Call the Router function from router.go
	router.Router(r)
	r.Run(os.Getenv("RUN_PORT"))
}
