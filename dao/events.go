package dao

import (
	"context"
	"fmt"

	"github.com/hello/databases"
	"github.com/hello/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEventById(id string) {
	e := models.Events{}
	fmt.Print(e)
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
