// File: internal/core/link_public.go
// Purpose: Core logic for resolving public slugs and handling shortlink access analytics.

package core

import (
	"context"
	"errors"
	"time"

	"shorty/internal/db"
	"shorty/internal/models"
)

var (
	ErrNotFound    = errors.New("short link not found")
	ErrLinkExpired = errors.New("link has expired")
	ErrLinkLocked  = errors.New("link requires password")
	ErrLinkBlocked = errors.New("link is disabled")
)

func ResolvePublicSlug(ctx context.Context, slug string, password string) (*models.Link, error) {
	var link models.Link
	err := db.FindOne(ctx, &link, "slug = ? AND disabled = false", slug)
	if err != nil {
		return nil, ErrNotFound
	}

	if link.ExpiresAt != nil && time.Now().After(*link.ExpiresAt) {
		return nil, ErrLinkExpired
	}

	if link.Password != "" && link.Password != password {
		return nil, ErrLinkLocked
	}

	go updateAccess(link.ID) // fire-and-forget analytics update

	return &link, nil
}

func updateAccess(id string) {
	ctx := context.Background()
	db.Exec(ctx, "UPDATE links SET click_count = click_count + 1, last_accessed = ? WHERE id = ?", time.Now(), id)
}
