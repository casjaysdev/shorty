// File: internal/core/link_org_actions.go
// Purpose: Core logic for handling org-scoped link operations (creation, update, deletion, analytics).

package core

import (
	"context"
	"errors"
	"time"

	"shorty/internal/db"
	"shorty/internal/models"
	"shorty/internal/util/slugger"
)

func CreateOrgLink(ctx context.Context, orgID string, link *models.Link) (*models.Link, error) {
	if !IsOrgAdmin(ctx, orgID) {
		return nil, errors.New("unauthorized")
	}

	if link.Slug == "" || !slugger.IsCustom(link.Slug) {
		link.Slug = slugger.Generate()
		link.CustomSlug = false
	} else {
		link.CustomSlug = true
	}

	link.ID = slugger.UUID()
	link.OrgID = orgID
	link.CreatedAt = time.Now()
	link.ClickCount = 0

	if err := db.Insert(ctx, link); err != nil {
		return nil, err
	}

	return link, nil
}

func ListOrgLinks(ctx context.Context, orgID string) ([]*models.Link, error) {
	if !IsOrgAdmin(ctx, orgID) {
		return nil, errors.New("unauthorized")
	}

	var links []*models.Link
	if err := db.FindWhere(ctx, &links, "org_id = ?", orgID); err != nil {
		return nil, err
	}
	return links, nil
}

func GetOrgLink(ctx context.Context, orgID string, id string) (*models.Link, error) {
	if !IsOrgAdmin(ctx, orgID) {
		return nil, errors.New("unauthorized")
	}

	var link models.Link
	if err := db.FindOne(ctx, &link, "id = ? AND org_id = ?", id, orgID); err != nil {
		return nil, err
	}
	return &link, nil
}

func UpdateOrgLink(ctx context.Context, orgID string, id string, updated *models.Link) (*models.Link, error) {
	link, err := GetOrgLink(ctx, orgID, id)
	if err != nil {
		return nil, err
	}

	link.OriginalURL = updated.OriginalURL
	link.Title = updated.Title
	link.Tags = updated.Tags
	link.Note = updated.Note
	link.Password = updated.Password
	link.ExpiresAt = updated.ExpiresAt
	link.OneTime = updated.OneTime
	link.Preview = updated.Preview
	link.Domain = updated.Domain

	if err := db.Update(ctx, link); err != nil {
		return nil, err
	}
	return link, nil
}

func DeleteOrgLink(ctx context.Context, orgID string, id string) error {
	link, err := GetOrgLink(ctx, orgID, id)
	if err != nil {
		return err
	}
	return db.Delete(ctx, link)
}

func ToggleOrgLink(ctx context.Context, orgID string, id string) (*models.Link, error) {
	link, err := GetOrgLink(ctx, orgID, id)
	if err != nil {
		return nil, err
	}
	link.Disabled = !link.Disabled
	if err := db.Update(ctx, link); err != nil {
		return nil, err
	}
	return link, nil
}

func CloneOrgLink(ctx context.Context, orgID string, id string) (*models.Link, error) {
	link, err := GetOrgLink(ctx, orgID, id)
	if err != nil {
		return nil, err
	}
	newLink := *link
	newLink.ID = slugger.UUID()
	newLink.Slug = slugger.Generate()
	newLink.CreatedAt = time.Now()
	newLink.ClickCount = 0
	if err := db.Insert(ctx, &newLink); err != nil {
		return nil, err
	}
	return &newLink, nil
}

func GetOrgLinkAnalytics(ctx context.Context, orgID string, id string) (map[string]interface{}, error) {
	link, err := GetOrgLink(ctx, orgID, id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"click_count":   link.ClickCount,
		"last_accessed": link.LastAccessed,
		"created_at":    link.CreatedAt,
		"tags":          link.Tags,
		"disabled":      link.Disabled,
	}, nil
}
