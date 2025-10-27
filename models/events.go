package models

import "time"

type Events struct {
	ID          string
	Name        string
	payload     map[string]interface{}
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
