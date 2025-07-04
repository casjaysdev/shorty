// File: internal/models/user.go
// Purpose: Defines the User model for authentication, settings, preferences, and dashboard data.

package models

import "time"

type User struct {
	ID             string     `json:"id" gorm:"primaryKey"`
	Username       string     `json:"username"`
	Email          string     `json:"email"`
	Password       string     `json:"-"` // stored hashed
	Theme          string     `json:"theme"` // dark, light, dracula
	IsAdmin        bool       `json:"is_admin"`
	OrgID          string     `json:"org_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	LastLogin      *time.Time `json:"last_login,omitempty"`
	RateLimit      int        `json:"rate_limit"`      // max links/day
	Plan           string     `json:"plan"`            // free, pro, business
	SlugDisabled   bool       `json:"slug_disabled"`   // disable custom slug
	SlugCase       string     `json:"slug_case"`       // mixed, lower, upper
	SlugLength     int        `json:"slug_length"`     // default length
	DomainOverride string     `json:"domain_override"` // optional default domain
}
