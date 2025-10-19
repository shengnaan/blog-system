package domain

import "github.com/google/uuid"

type User struct {
	Base
	Email       string      `json:"email"`
	PassHash    string      `json:"-"`
	DisplayName string      `json:"display_name"`
	AvatarURL   *string     `json:"avatar_url,omitempty"`
	IsBlocked   bool        `json:"is_blocked"`
	Roles       []Role      `json:"roles,omitempty"`
	Blogs       []uuid.UUID `json:"blogs,omitempty"`
}
