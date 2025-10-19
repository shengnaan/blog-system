package domain

import (
	"time"

	"github.com/google/uuid"
)

type PostStatus string

const (
	PostDraft     PostStatus = "draft"
	PostPending   PostStatus = "pending"
	PostPublished PostStatus = "published"
	PostHidden    PostStatus = "hidden"
	PostArchived  PostStatus = "archived"
)

type Post struct {
	Base
	BlogID      uuid.UUID  `json:"blog_id"`
	OwnerID     uuid.UUID  `json:"owner_id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	ContentMD   string     `json:"content_md"`
	ContentHTML *string    `json:"content_html,omitempty"`
	Status      PostStatus `json:"status"`
	Views       int        `json:"views"`
	Likes       int        `json:"likes"`
	PublishedAt *time.Time `json:"published_at,omitempty"`

	Tags []Tag `json:"tags,omitempty"`
}
