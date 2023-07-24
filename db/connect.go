package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	UserColl = DB.Collection("users")
	TeamColl = DB.Collection("teams")
)

var (
	DB = connectDB()
)

func connectDB() *mongo.Database {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Error(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("successfully connecting to database")
	return client.Database(dbName)
}
