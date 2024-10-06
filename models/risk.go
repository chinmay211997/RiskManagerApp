package models

import "github.com/google/uuid"

type Risk struct {
	ID          uuid.UUID `json:"id"`
	State       RiskState `json:"state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type RiskState string

const (
	OPEN          RiskState = "open"
	CLOSED        RiskState = "closed"
	ACCEPTED      RiskState = "accepted"
	INVESTIGATING RiskState = "investigating"
)
