// File: internal/storage/gorm/gorm.go
// Purpose: GORM-backed implementation of the Storage interface for Shorty.

package gormstorage

import (
	"context"
	"time"

	"gorm.io/gorm"

	"shorty/internal/storage"
)

type GormStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) CreateLink(ctx context.Context, link *storage.Link) error {
	link.CreatedAt = time.Now()
	return s.db.WithContext(ctx).Create(link).Error
}

func (s *GormStore) GetLinkBySlug(ctx context.Context, slug string) (*storage.Link, error) {
	var link storage.Link
	err := s.db.WithContext(ctx).Where("slug = ?", slug).First(&link).Error
	return &link, err
}

func (s *GormStore) GetLinksByUser(ctx context.Context, userID string) ([]*storage.Link, error) {
	var links []*storage.Link
	err := s.db.WithContext(ctx).Where("created_by = ?", userID).Find(&links).Error
	return links, err
}

func (s *GormStore) UpdateLink(ctx context.Context, link *storage.Link) error {
	return s.db.WithContext(ctx).Save(link).Error
}

func (s *GormStore) DeleteLink(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&storage.Link{}, "id = ?", id).Error
}

func (s *GormStore) IncrementClick(ctx context.Context, slug string) error {
	return s.db.WithContext(ctx).Model(&storage.Link{}).Where("slug = ?", slug).UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error
}

func (s *GormStore) SearchLinks(ctx context.Context, query string) ([]*storage.Link, error) {
	var links []*storage.Link
	err := s.db.WithContext(ctx).Where("slug LIKE ? OR url LIKE ?", "%"+query+"%", "%"+query+"%").Find(&links).Error
	return links, err
}
