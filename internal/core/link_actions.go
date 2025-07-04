// File: internal/core/link_actions.go
// Purpose: Core logic for handling link creation, resolution, deletion, and disabling in Shorty.

package core

import (
	"errors"
	"time"

	"shorty/internal/lib/utils"
	"shorty/internal/models"
	"shorty/internal/storage"
)

// CreateLink handles full logic for creating a shortened link.
func CreateLink(link *models.Link) (*models.Link, error) {
	if link.OriginalURL == "" {
		return nil, errors.New("original URL is required")
	}
	if link.Slug == "" {
		link.Slug = utils.GenerateSlug()
	}
	link.CreatedAt = time.Now()
	link.ClickCount = 0

	if err := storage.DB().Create(link).Error; err != nil {
		return nil, err
	}
	return link, nil
}

// ResolveLink returns the original URL if the slug exists and is active.
func ResolveLink(slug string) (*models.Link, error) {
	var link models.Link
	if err := storage.DB().Where("slug = ? AND disabled = false", slug).First(&link).Error; err != nil {
		return nil, errors.New("link not found or disabled")
	}

	now := time.Now()
	link.ClickCount++
	link.LastAccessed = &now
	storage.DB().Save(&link)

	// One-time link logic
	if link.OneTime {
		link.Disabled = true
		storage.DB().Save(&link)
	}

	return &link, nil
}

// DisableLink disables a link by slug or ID.
func DisableLink(slugOrID string) error {
	return storage.DB().Model(&models.Link{}).
		Where("slug = ? OR id = ?", slugOrID, slugOrID).
		Update("disabled", true).Error
}

// DeleteLink removes a link from storage.
func DeleteLink(slugOrID string) error {
	return storage.DB().Where("slug = ? OR id = ?", slugOrID, slugOrID).Delete(&models.Link{}).Error
}

// FindLink returns a link by slug or ID without modifying it.
func FindLink(slugOrID string) (*models.Link, error) {
	var link models.Link
	if err := storage.DB().Where("slug = ? OR id = ?", slugOrID, slugOrID).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// SearchLinks returns matching links based on partial slug, tag, or URL.
func SearchLinks(query string) ([]models.Link, error) {
	var links []models.Link
	err := storage.DB().Where("slug LIKE ? OR original_url LIKE ? OR tags LIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Find(&links).Error
	return links, err
}
