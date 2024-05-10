package models

import (
	"github.com/google/uuid"
)

type Video struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	VideoID      string    `gorm:"type:varchar(100);not null" json:"video_id"`
	Title        string    `gorm:"type:varchar(255);not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	PublishedAt  string    `gorm:"type:timestamp;not null" json:"published_at"`
	ThumbnailURL string    `gorm:"type:varchar(255);not null" json:"thumbnail_url"`
}
