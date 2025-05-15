package services

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/moinuddinshovan/url-shortener-go/internal/models"
	"gorm.io/gorm"
)

// URLService handles business logic for URL operations
type URLService struct {
	db *gorm.DB
}

// NewURLService creates a new URLService instance
func NewURLService(db *gorm.DB) *URLService {
	return &URLService{db: db}
}

// ValidateURL checks if the given URL is valid
func (s *URLService) ValidateURL(rawURL string) (string, error) {
	// Add https:// if no protocol is specified
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	// Parse the URL to ensure it's valid
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.New("invalid URL provided")
	}

	// Check if the URL has a valid host
	if parsedURL.Host == "" {
		return "", errors.New("invalid URL: missing host")
	}

	// Ensure the URL has a scheme
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}

	// Reconstruct the URL to ensure it's properly formatted
	return parsedURL.String(), nil
}

// CreateShortURL creates a new shortened URL
func (s *URLService) CreateShortURL(originalURL string) (*models.URL, error) {
	validatedURL, err := s.ValidateURL(originalURL)
	if err != nil {
		return nil, err
	}

	// Generate a unique short code
	shortCode := uuid.New().String()[:8]

	url := &models.URL{
		OriginalURL: validatedURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}

	if err := s.db.Create(url).Error; err != nil {
		return nil, err
	}

	return url, nil
}

// GetURLByShortCode retrieves a URL by its short code
func (s *URLService) GetURLByShortCode(shortCode string) (*models.URL, error) {
	var url models.URL
	if err := s.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}
	return &url, nil
}

// GetAllURLs retrieves all URLs
func (s *URLService) GetAllURLs() ([]models.URL, error) {
	var urls []models.URL
	if err := s.db.Order("created_at desc").Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

// IncrementClicks increments the click count for a URL
func (s *URLService) IncrementClicks(url *models.URL) error {
	url.Clicks++
	return s.db.Save(url).Error
}

// UpdateURL updates an existing URL
func (s *URLService) UpdateURL(url *models.URL, newURL string) error {
	validatedURL, err := s.ValidateURL(newURL)
	if err != nil {
		return err
	}

	url.OriginalURL = validatedURL
	return s.db.Save(url).Error
}

// DeleteURL deletes a URL
func (s *URLService) DeleteURL(url *models.URL) error {
	return s.db.Delete(url).Error
}

// GetDatabaseContents retrieves raw database contents
func (s *URLService) GetDatabaseContents() (map[string]interface{}, error) {
	var urls []models.URL
	if err := s.db.Find(&urls).Error; err != nil {
		return nil, errors.New("failed to retrieve database contents")
	}

	// Convert URLs to a more detailed format
	urlMaps := make([]map[string]interface{}, len(urls))
	for i, url := range urls {
		urlMaps[i] = map[string]interface{}{
			"id":           url.ID,
			"original_url": url.OriginalURL,
			"short_code":   url.ShortCode,
			"clicks":       url.Clicks,
			"created_at":   url.CreatedAt,
		}
	}

	return map[string]interface{}{
		"total_records": len(urls),
		"urls":          urlMaps,
	}, nil
}
