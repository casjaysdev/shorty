// File: internal/models/domain_ban.go
// Purpose: Stores banned slugs/domains and reasons for administrative enforcement.

package models

import (
	"time"
)

type DomainBan struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Target    string    `json:"target"` // slug or domain
	Reason    string    `json:"reason,omitempty"`
	AdminNote string    `json:"admin_note,omitempty"`
	PredefinedReason string `json:"predefined_reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IsDomain  bool      `json:"is_domain"` // false means it's a slug
	CreatedBy string    `json:"created_by"`
}
