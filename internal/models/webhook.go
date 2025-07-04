// File: internal/models/webhook.go
// Purpose: Represents a webhook configuration including provider, scope, and retry state.

package models

import "time"

type Webhook struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	OwnerID    string    `json:"owner_id"`     // user/org/system
	Scope      string    `json:"scope"`        // user | org | server
	Provider   string    `json:"provider"`     // slack, discord, zapier, etc
	URL        string    `json:"url"`
	Secret     string    `json:"secret"`       // for HMAC signing
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastUsed   *time.Time `json:"last_used,omitempty"`
	EventTypes string    `json:"event_types"`  // comma-separated events (paste.created, paste.deleted, etc)
	RetryCount int       `json:"retry_count"`
}
