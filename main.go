package main

import (
	"chatroom-api/database"
	"chatroom-api/entities"
	"chatroom-api/router"
	"os"

	"github.com/gin-contrib/cors"
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

	// Use CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to frontend URL in production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "x-api-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Call the Router function from router.go
	router.Router(r)
	r.Run(os.Getenv("RUN_PORT"))
}
