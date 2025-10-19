package domain

import "github.com/google/uuid"

type Tag struct {
	Base
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PostTag struct {
	Base
	PostID uuid.UUID `json:"post_id"`
	TagID  uuid.UUID `json:"tag_id"`
}
