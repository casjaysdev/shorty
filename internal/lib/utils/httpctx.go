// File: internal/lib/utils/httpctx.go
// Purpose: Utilities for working with HTTP context, extracting scoped user/org/session data.

package utils

import (
	"context"
	"net/http"
)

type ctxKey string

const (
	CtxUserID  ctxKey = "user_id"
	CtxOrgID   ctxKey = "org_id"
	CtxTokenID ctxKey = "token_id"
)

// GetUserID extracts the user ID from request context.
func GetUserID(r *http.Request) string {
	val := r.Context().Value(CtxUserID)
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// GetOrgID extracts the org ID from request context.
func GetOrgID(r *http.Request) string {
	val := r.Context().Value(CtxOrgID)
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// GetTokenID extracts the token ID from request context.
func GetTokenID(r *http.Request) string {
	val := r.Context().Value(CtxTokenID)
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// InjectContext adds user/org/token IDs to the context.
func InjectContext(r *http.Request, userID, orgID, tokenID string) *http.Request {
	ctx := context.WithValue(r.Context(), CtxUserID, userID)
	ctx = context.WithValue(ctx, CtxOrgID, orgID)
	ctx = context.WithValue(ctx, CtxTokenID, tokenID)
	return r.WithContext(ctx)
}
