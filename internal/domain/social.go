package domain

import "github.com/google/uuid"

type PostLike struct {
	Base
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

type Bookmark struct {
	Base
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}
