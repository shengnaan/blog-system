package domain

import "github.com/google/uuid"

type Blog struct {
	Base
	OwnerID     uuid.UUID `json:"owner_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	IsPublic    bool      `json:"is_public"`
}
