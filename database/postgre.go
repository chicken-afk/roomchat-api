package database

import (
	"fmt"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	/* Load ENV */
	errEnv := godotenv.Load()
	if errEnv != nil {
		logrus.Fatal("LOAD ENV ", errEnv)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbInstance := os.Getenv("DB_INSTANCE")

	var dialector gorm.Dialector

	if os.Getenv("SERVICE_MODE") == "develop" {
		// Use standard PostgreSQL driver for local development
		logrus.Info("Using local development mode")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort)
		dialector = postgres.Open(dsn)
	} else {
		// Use Cloud SQL Proxy for production mode
		logrus.Info("Using cloudsqlproxy production mode")
		dialector = postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN: fmt.Sprintf("instance=%s user=%s dbname=%s password=%s sslmode=disable",
				dbInstance, dbUser, dbName, dbPass),
		})
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database: ", err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSql.Close()
}
