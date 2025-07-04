// File: internal/models/org.go
// Purpose: Defines the Organization model, including metadata, billing, custom domains, and plan enforcement.

package models

import "time"

type Org struct {
	ID             string     `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	OwnerID        string     `json:"owner_id"` // User who created the org
	CreatedAt      time.Time  `json:"created_at"`
	BillingEmail   string     `json:"billing_email"`
	Plan           string     `json:"plan"` // free, pro, business
	TrialExpiresAt *time.Time `json:"trial_expires_at,omitempty"`
	SlugPolicy     string     `json:"slug_policy"` // allow, deny, warn
	CustomDomain   string     `json:"custom_domain,omitempty"`
	EmailSender    string     `json:"email_sender,omitempty"` // FROM: name
	SMTPServer     string     `json:"smtp_server,omitempty"`
	SMTPUser       string     `json:"smtp_user,omitempty"`
	SMTPPassword   string     `json:"-"` // encrypted or hidden
	Theme          string     `json:"theme"` // optional org-level theme
	RateLimit      int        `json:"rate_limit"`
	AllowWhiteLabel bool      `json:"allow_white_label"`
}
