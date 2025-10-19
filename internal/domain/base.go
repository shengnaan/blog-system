package domain

import (
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

var UndeletedValue = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewBase() Base {
	now := time.Now().UTC()
	return Base{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: UndeletedValue,
	}
}

func (b *Base) Touch() { b.UpdatedAt = time.Now().UTC() }

func (b *Base) SoftDelete() { b.DeletedAt = time.Now().UTC() }

func (b *Base) Undelete() { b.DeletedAt = UndeletedValue }

func (b *Base) IsDeleted() bool { return !b.DeletedAt.Equal(UndeletedValue) }

func TableName[T any]() string {
	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := strings.ToLower(t.Name())
	if strings.HasSuffix(name, "y") &&
		!strings.HasSuffix(name, "ay") &&
		!strings.HasSuffix(name, "ey") &&
		!strings.HasSuffix(name, "iy") &&
		!strings.HasSuffix(name, "oy") &&
		!strings.HasSuffix(name, "uy") {
		return name[:len(name)-1] + "ies"
	}
	if strings.HasSuffix(name, "s") {
		return name + "es"
	}
	return name + "s"
}
