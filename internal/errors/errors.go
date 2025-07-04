// File: internal/errors/errors.go
// Purpose: Defines reusable error variables and helpers for the Shorty service.

package errors

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrConflict     = errors.New("conflict")
	ErrInvalidSlug  = errors.New("invalid slug format")
	ErrSlugTaken    = errors.New("slug already in use")
	ErrLinkInactive = errors.New("link is no longer active")
	ErrRateLimited  = errors.New("rate limit exceeded")
	ErrBannedSlug   = errors.New("slug has been banned")
	ErrBannedDomain = errors.New("domain has been banned")
	ErrMissingInput = errors.New("required input missing")
	ErrInvalidToken = errors.New("invalid or expired token")
)
