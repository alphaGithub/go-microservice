package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hello/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoCollection = make(map[string]*mongo.Collection)
var DatabaseCollections = []string{"events"}

func connectMongoDB(connectionURL string, dbName string, collectionName []string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("[err] Mongo connection failed >>>>>", err)
		return nil, err
	}

	// ✅ Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("[err] Mongo ping failed >>>>> %v", err)
		return nil, err
	}
	fmt.Println("[msg] MongoDB connected successfully.😊")
	MongoClient = client

	for i := range len(collectionName) {
		cName := collectionName[i]
		collection := client.Database(dbName).Collection(cName)
		fmt.Println("[test] >>>>> ", cName, dbName, collection)
		MongoCollection[cName] = collection
	}
	return client, err
}

func InitDb() {
	fmt.Println("[msg] database initializing ....")
	connectionURL := config.Config.Databases.Hello.ConnectionURL
	databaseName := config.Config.Databases.Hello.DbName
	_, err := connectMongoDB(connectionURL, databaseName, DatabaseCollections)

	if err != nil {
		fmt.Print("[err] Panic >>>> database connect failed")
		panic(err)
	}
}
