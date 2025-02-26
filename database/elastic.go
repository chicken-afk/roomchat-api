package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func SetupElasticConnection() *elastic.Client {
	// Setup ElasticSearch Connection
	errEnv := godotenv.Load()
	if errEnv != nil {
		logrus.Fatal("LOAD ENV ", errEnv)
	}

	// Load ENV
	elasticHost := os.Getenv("ELASTICSEARCH_HOST")
	elasticUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	elasticPassword := os.Getenv("ELASTICSEARCH_PASSWORD")
	logrus.Info("ELASTIC_HOST: ", elasticHost)
	logrus.Info("ELASTIC_USERNAME: ", elasticUsername)
	logrus.Info("ELASTIC_PASSWORD: ", elasticPassword)

	// Connect to ElasticSearch
	elasticClient, err := elastic.NewClient(
		elastic.SetURL(elasticHost),
		elastic.SetBasicAuth(elasticUsername, elasticPassword),
		elastic.SetSniff(false),
	)
	if err != nil {
		logrus.Fatal("Failed to connect to ElasticSearch: ", err)
	}
	return elasticClient
}
