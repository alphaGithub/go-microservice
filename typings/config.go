package typings

import "time"

type CreateEventPayloadType struct {
	Name        string
	Description string
	Payload     map[string]interface{}
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
