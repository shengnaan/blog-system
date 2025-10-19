package domain

import "github.com/google/uuid"

type Role string

const (
	RoleUser      Role = "user"
	RoleAuthor    Role = "author"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
)

type UserRole struct {
	Base
	UserID   uuid.UUID `json:"user_id"`
	RoleCode Role      `json:"role_code"`
}
