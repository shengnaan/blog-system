package domain

import "github.com/google/uuid"

type CommentStatus string

const (
	CommentVisible CommentStatus = "visible"
	CommentHidden  CommentStatus = "hidden"
	CommentDeleted CommentStatus = "deleted"
)

type Comment struct {
	Base
	PostID   uuid.UUID     `json:"post_id"`
	AuthorID uuid.UUID     `json:"author_id"`
	ParentID *uuid.UUID    `json:"parent_id,omitempty"`
	Body     string        `json:"body"`
	Status   CommentStatus `json:"status"`
}
