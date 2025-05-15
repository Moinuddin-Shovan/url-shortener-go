package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moinuddinshovan/url-shortener-go/internal/services"
)

// URLHandler handles HTTP requests for URL operations
type URLHandler struct {
	urlService *services.URLService
}

// NewURLHandler creates a new URLHandler instance
func NewURLHandler(urlService *services.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

// RegisterRoutes registers all URL-related routes
func (h *URLHandler) RegisterRoutes(router *gin.Engine) {
	// Web routes
	router.GET("/", h.renderIndex)
	router.LoadHTMLGlob("web/templates/*")

	// API routes
	api := router.Group("/api/urls")
	{
		api.GET("", h.getAllURLs)
		api.POST("", h.createShortURL)
		api.GET("/:shortCode", h.getURL)
		api.PUT("/:shortCode", h.updateURL)
		api.DELETE("/:shortCode", h.deleteURL)
		api.GET("/:shortCode/stats", h.getURLStats)
		api.GET("/db/contents", h.getDatabaseContents)
	}

	// Redirect route
	router.GET("/:shortCode", h.redirectToURL)
}

// renderIndex renders the main page
func (h *URLHandler) renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// getAllURLs handles GET /api/urls
func (h *URLHandler) getAllURLs(c *gin.Context) {
	urls, err := h.urlService.GetAllURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	urlMaps := make([]map[string]interface{}, len(urls))
	for i, url := range urls {
		urlMaps[i] = url.ToMap()
	}

	c.JSON(http.StatusOK, urlMaps)
}

// createShortURL handles POST /api/urls
func (h *URLHandler) createShortURL(c *gin.Context) {
	var input struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	url, err := h.urlService.CreateShortURL(input.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url":    c.Request.Host + "/" + url.ShortCode,
		"original_url": url.OriginalURL,
		"short_code":   url.ShortCode,
	})
}

// getURL handles GET /api/urls/:shortCode
func (h *URLHandler) getURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := h.urlService.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, url.ToMap())
}

// updateURL handles PUT /api/urls/:shortCode
func (h *URLHandler) updateURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := h.urlService.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	var input struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	if err := h.urlService.UpdateURL(url, input.URL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, url.ToMap())
}

// deleteURL handles DELETE /api/urls/:shortCode
func (h *URLHandler) deleteURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := h.urlService.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if err := h.urlService.DeleteURL(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete URL"})
		return
	}

	c.Status(http.StatusNoContent)
}

// getURLStats handles GET /api/urls/:shortCode/stats
func (h *URLHandler) getURLStats(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := h.urlService.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_code":   url.ShortCode,
		"original_url": url.OriginalURL,
		"clicks":       url.Clicks,
		"created_at":   url.CreatedAt,
	})
}

// getDatabaseContents handles GET /api/urls/db/contents
func (h *URLHandler) getDatabaseContents(c *gin.Context) {
	contents, err := h.urlService.GetDatabaseContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve database contents"})
		return
	}

	c.JSON(http.StatusOK, contents)
}

// redirectToURL handles GET /:shortCode
func (h *URLHandler) redirectToURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := h.urlService.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if err := h.urlService.IncrementClicks(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update click count"})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}
