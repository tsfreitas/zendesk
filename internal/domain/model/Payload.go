package model

import "encoding/json"

type Payload struct {
	Event   Event           `json:"event_type"`
	Payload json.RawMessage `json:"payload"`
}

type Event string

const (
	CreateUserAndOrganization Event = "CREATE_USER_AND_ORGANIZATION"
)
