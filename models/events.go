package models

import (
	"time"
)

type Event struct {
	ID          string                 `bson:"_id,omitempty"  json:"id"`
	ShortId     string                 `bson:"short_id" json:"short_id"`
	Name        string                 `bson:"name" json:"name"`
	MetaData    map[string]interface{} `bson:"payload" json:"meta_data"`
	Description string                 `bson:"description" json:"description"`
	CreatedAt   time.Time              `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at" json:"updated_at"`
}
