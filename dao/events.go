package dao

import (
	"context"
	"fmt"

	"github.com/hello/databases"
	"github.com/hello/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEventById(id string) (models.Event, error) {
	collection := databases.MongoCollection["Event"]
	var event models.Event
	filter := bson.M{"short_id": id}
	result := collection.FindOne(context.Background(), filter)
	if err := result.Decode(&event); err != nil {
		return event, err
	}
	return event, nil
}

func GetEvents(limit int, offset int) ([]models.Event, error) {
	queryOptions := options.Find()
	queryOptions.SetLimit(int64(limit))
	queryOptions.SetSkip(int64(offset))
	collection := databases.MongoCollection["events"]
	fmt.Println("queryOptions:", queryOptions)
	fmt.Println("collection:", collection)
	cursor, err := collection.Find(context.Background(), bson.D{}, queryOptions)

	if err != nil {
		fmt.Println("[err] >>>>> [GetEvent|dao] >>>>> failed to fetch from database", err)
		return []models.Event{}, err
	}
	var result []models.Event
	err = cursor.All(context.Background(), &result)
	if err != nil {
		fmt.Println("[err] >>>>> [GetEvent|dao] error with cursor", err)
		return []models.Event{}, err
	}
	fmt.Println(result)
	defer cursor.Close(context.Background())
	return result, nil
}

func CreateEvent(event *models.Event) (interface{}, error) {
	collection := databases.MongoCollection["Event"]
	result, err := collection.InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, err
}
