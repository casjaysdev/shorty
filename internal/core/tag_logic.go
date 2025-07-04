// File: internal/core/tag_logic.go
// Purpose: Provides core functions for tag filtering and aggregation used by search, filters, and suggestions.

package core

import (
	"strings"

	"shorty/internal/models"
	"shorty/internal/storage"
)

// GetAllTags retrieves a unique, sorted list of all tags in use.
func GetAllTags() ([]string, error) {
	var tags []string
	var results []models.Link

	if err := storage.DB().Select("tags").Find(&results).Error; err != nil {
		return nil, err
	}

	tagSet := make(map[string]struct{})
	for _, link := range results {
		if link.Tags == "" {
			continue
		}
		for _, tag := range strings.Split(link.Tags, ",") {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				tagSet[tag] = struct{}{}
			}
		}
	}

	for tag := range tagSet {
		tags = append(tags, tag)
	}
	return tags, nil
}

// FilterLinksByTag returns links that match the given tag.
func FilterLinksByTag(tag string) ([]models.Link, error) {
	var links []models.Link
	if tag == "" {
		return links, nil
	}
	err := storage.DB().
		Where("tags LIKE ?", "%"+tag+"%").
		Find(&links).Error
	return links, err
}
