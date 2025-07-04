// File: internal/storage/storage.go
// Purpose: Defines the interface for interacting with persistent short link storage.

package storage

import (
	"context"
	"time"
)

type Link struct {
	ID         string
	Slug       string
	URL        string
	CreatedAt  time.Time
	CreatedBy  string
	AccessedAt time.Time
	Clicks     int
	Active     bool
	Domain     string
}

type Storage interface {
	CreateLink(ctx context.Context, link *Link) error
	GetLinkBySlug(ctx context.Context, slug string) (*Link, error)
	GetLinksByUser(ctx context.Context, userID string) ([]*Link, error)
	UpdateLink(ctx context.Context, link *Link) error
	DeleteLink(ctx context.Context, id string) error
	IncrementClick(ctx context.Context, slug string) error
	SearchLinks(ctx context.Context, query string) ([]*Link, error)
}
