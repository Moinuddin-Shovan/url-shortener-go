package models

import (
	"time"

	"gorm.io/gorm"
)

// URL represents a shortened URL in the database
type URL struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	OriginalURL string    `gorm:"not null" json:"original_url"`
	ShortCode   string    `gorm:"uniqueIndex;not null" json:"short_code"`
	Clicks      int       `gorm:"default:0" json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BeforeCreate is a GORM hook that generates a short code before creating a new URL
func (u *URL) BeforeCreate(tx *gorm.DB) error {
	// Generate a short code using the first 8 characters of a UUID
	u.ShortCode = u.GenerateShortCode()
	return nil
}

// GenerateShortCode generates a unique 8-character code
func (u *URL) GenerateShortCode() string {
	// This is a simple implementation. In production, you might want to use
	// a more sophisticated algorithm to ensure uniqueness and avoid collisions
	return u.ShortCode[:8]
}

// ToMap converts the URL struct to a map for JSON response
func (u *URL) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           u.ID,
		"original_url": u.OriginalURL,
		"short_code":   u.ShortCode,
		"clicks":       u.Clicks,
		"created_at":   u.CreatedAt,
		"updated_at":   u.UpdatedAt,
	}
}
