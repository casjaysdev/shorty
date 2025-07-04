// File: internal/models/link.go
// Purpose: Defines the Link model used by Shorty for URL storage, API, and persistence.

package models

import (
	"time"
)

type Link struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	UserID       string     `json:"user_id"`
	OrgID        string     `json:"org_id,omitempty"`
	OriginalURL  string     `json:"original_url"`
	Slug         string     `json:"slug" gorm:"uniqueIndex"`
	CustomSlug   bool       `json:"custom_slug"`
	Password     string     `json:"-"` // stored as hash if used
	CreatedAt    time.Time  `json:"created_at"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	OneTime      bool       `json:"one_time"`
	Disabled     bool       `json:"disabled"`
	ClickCount   int        `json:"click_count"`
	LastAccessed *time.Time `json:"last_accessed,omitempty"`
	Tags         string     `json:"tags,omitempty"` // comma-separated
	Domain       string     `json:"domain,omitempty"`
}
