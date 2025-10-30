package dao

import (
	"context"
	"fmt"

	"github.com/hello/databases"
	"github.com/hello/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEventById(id string) (models.Events, error) {
	collection := databases.MongoCollection["events"]
	var event models.Events
	filter := bson.M{"short_id": id}
	result := collection.FindOne(context.Background(), filter)
	if err := result.Decode(&event); err != nil {
		return event, err
	}
	return event, nil
}

func GetEvents(limit int, offset int) ([]models.Events, error) {
	queryOptions := options.Find()
	queryOptions.SetLimit(int64(limit))
	queryOptions.SetSkip(int64(offset))
	collection := databases.MongoCollection["events"]
	cursor, err := collection.Find(context.Background(), bson.D{}, queryOptions)

	if err != nil {
		fmt.Println("[err] >>>>> [GetEvents|dao] >>>>> failed to fetch from database", err)
		return []models.Events{}, err
	}
	var result []models.Events
	err = cursor.All(context.Background(), &result)
	if err != nil {
		fmt.Println("[err] >>>>> [GetEvents|dao] error with cursor", err)
		return []models.Events{}, err
	}
	fmt.Println(result)
	defer cursor.Close(context.Background())
	return result, nil
}

func CreateEvents(event *models.Events) (interface{}, error) {
	collection := databases.MongoCollection["events"]
	result, err := collection.InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, err
}
