package models

type Pin struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description" validate:"max=200"`
}
