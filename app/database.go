package app

import (
	"bwastartupecho/helper"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName                  string = "bwastartup"
	userCollection          string = "users"
	campaignCollection      string = "campaigns"
	transactionCOllection   string = "transactions"
	campaignImageCollection string = "campaign_images"
)

type MongoDatabase interface {
	CloseDB()
	DbName() *mongo.Database
	UserCollection() *mongo.Collection
	CampaignImageCollection() *mongo.Collection
	CampaignCollection() *mongo.Collection
	TransactionCollection() *mongo.Collection
}

type dbstruct struct {
	db *mongo.Client
}

// Connection URI

func NewDB() MongoDatabase {
	const uri = "mongodb://127.0.0.1:27017/?maxPoolSize=20&w=1"
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	helper.PanicIfError(err, "error in connect to mongo server")
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Ping the primary
	err = client.Ping(context.TODO(), readpref.Primary())
	helper.PanicIfError(err, " error in ping connection")
	// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Successfully connected and pinged.")

	return &dbstruct{client}

}

func (db *dbstruct) CloseDB() {
	err := db.db.Disconnect(context.TODO())
	helper.PanicIfError(err, " erro in close connection mongo")
	log.Println("db is already closed")
}

func (db *dbstruct) DbName() *mongo.Database {
	database := db.db.Database(dbName)
	return database
}

func (db *dbstruct) UserCollection() *mongo.Collection {
	userCOllection := db.DbName().Collection(userCollection)
	return userCOllection
}

func (db *dbstruct) CampaignCollection() *mongo.Collection {
	return db.DbName().Collection(campaignCollection)
}

func (db *dbstruct) CampaignImageCollection() *mongo.Collection {
	return db.DbName().Collection(campaignImageCollection)
}

func (db *dbstruct) TransactionCollection() *mongo.Collection {
	return db.DbName().Collection(transactionCOllection)
}
