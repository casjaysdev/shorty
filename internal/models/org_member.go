// File: internal/models/org_member.go
// Purpose: Defines the relationship between users and organizations, including roles, invites, and ownership.

package models

import "time"

type OrgMember struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	OrgID     string    `json:"org_id"`
	Role      string    `json:"role"` // admin, member
	JoinedAt  time.Time `json:"joined_at"`
	Invite    bool      `json:"invite"`
	Accepted  bool      `json:"accepted"`
}
