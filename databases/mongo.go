package databases

import (
	"log"

	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	"github.com/SwaDeshiTech/kubesync/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

func InitializeMongoConnection() error {

	mongoConfig := v1.Mongo{
		URI:      config.GetConfig().Mongo.URI,
		Username: config.GetConfig().Mongo.Username,
		Password: config.GetConfig().Mongo.Password,
		ConnectionPoolDetail: v1.ConnectionPoolDetail{
			MaxPoolSize:     config.GetConfig().Mongo.MaxPoolSize,
			MinPoolSize:     config.GetConfig().Mongo.MinPoolSize,
			MaxIdleTime:     config.GetConfig().Mongo.MaxIdleTime,
			MaxConnIdleTime: config.GetConfig().Mongo.MaxConnIdleTime,
			ConnectTimeout:  config.GetConfig().Mongo.ConnectTimeout,
		},
	}

	client, err := mongoConfig.InitializeMongoConnection()
	if err != nil {
		log.Println("Error in getting mongo client", err)
		return err
	}

	MongoClient = client

	return nil
}
