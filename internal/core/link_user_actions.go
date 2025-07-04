// File: internal/core/link_user_actions.go
// Purpose: Core logic for handling user-specific link operations.

package core

import (
	"context"
	"errors"
	"time"

	"shorty/internal/db"
	"shorty/internal/models"
	"shorty/internal/util/slugger"
)

func CreateLink(ctx context.Context, link *models.Link) (*models.Link, error) {
	userID := ContextUserID(ctx)
	if userID == "" {
		return nil, errors.New("unauthorized")
	}

	if link.Slug == "" || !slugger.IsCustom(link.Slug) {
		link.Slug = slugger.Generate()
		link.CustomSlug = false
	} else {
		link.CustomSlug = true
	}

	link.ID = slugger.UUID()
	link.UserID = userID
	link.CreatedAt = time.Now()
	link.ClickCount = 0

	if err := db.Insert(ctx, link); err != nil {
		return nil, err
	}

	return link, nil
}

func ListUserLinks(ctx context.Context) ([]*models.Link, error) {
	userID := ContextUserID(ctx)
	if userID == "" {
		return nil, errors.New("unauthorized")
	}

	var links []*models.Link
	if err := db.FindWhere(ctx, &links, "user_id = ?", userID); err != nil {
		return nil, err
	}
	return links, nil
}

func GetUserLink(ctx context.Context, id string) (*models.Link, error) {
	userID := ContextUserID(ctx)
	if userID == "" {
		return nil, errors.New("unauthorized")
	}

	var link models.Link
	if err := db.FindOne(ctx, &link, "id = ? AND user_id = ?", id, userID); err != nil {
		return nil, err
	}
	return &link, nil
}

func UpdateUserLink(ctx context.Context, id string, updated *models.Link) (*models.Link, error) {
	link, err := GetUserLink(ctx, id)
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

func DeleteUserLink(ctx context.Context, id string) error {
	link, err := GetUserLink(ctx, id)
	if err != nil {
		return err
	}
	return db.Delete(ctx, link)
}

func ToggleLink(ctx context.Context, id string) (*models.Link, error) {
	link, err := GetUserLink(ctx, id)
	if err != nil {
		return nil, err
	}
	link.Disabled = !link.Disabled
	if err := db.Update(ctx, link); err != nil {
		return nil, err
	}
	return link, nil
}

func CloneLink(ctx context.Context, id string) (*models.Link, error) {
	orig, err := GetUserLink(ctx, id)
	if err != nil {
		return nil, err
	}

	newLink := *orig
	newLink.ID = slugger.UUID()
	newLink.Slug = slugger.Generate()
	newLink.CreatedAt = time.Now()
	newLink.ClickCount = 0
	if err := db.Insert(ctx, &newLink); err != nil {
		return nil, err
	}
	return &newLink, nil
}

func GetLinkAnalytics(ctx context.Context, id string) (map[string]interface{}, error) {
	link, err := GetUserLink(ctx, id)
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
