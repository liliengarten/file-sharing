package models

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"max=200"`
	Private     bool   `json:"private" validate:""`
}

type BoardAuthor struct {
	UserID    int `json:"user_id"`
	InvitedBy int `json:"invited_by"`
}
