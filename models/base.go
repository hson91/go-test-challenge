package models

// Base : truct
type Base struct {
	ID         uint   `json:"_id"`
	ExternalID string `json:"external_id"`
	CreatedAt  string `json:"created_at"`
}
