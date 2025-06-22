# URL Shortener Service Documentation

## Overview
The URL Shortener is a Go-based web application that provides a service to convert long URLs into shorter, more manageable links. The application is built using the Gin web framework and uses SQLite for data persistence.

## Architecture

### 1. Project Structure
```
url-shortener-go/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── handlers/         # HTTP request handlers
│   │   └── url_handler.go
│   ├── models/          # Data models
│   │   └── url.go
│   └── services/        # Business logic
│       └── url_service.go
├── web/
│   └── templates/       # HTML templates
│       └── index.html
└── go.mod              # Go module file
```

### 2. Core Components

#### Models (`internal/models/url.go`)
The URL model represents the structure of our shortened URLs:
```go
type URL struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    OriginalURL string    `json:"original_url" gorm:"not null"`
    ShortCode   string    `json:"short_code" gorm:"unique;not null"`
    CreatedAt   time.Time `json:"created_at"`
    AccessCount int       `json:"access_count" gorm:"default:0"`
}
```

#### Services (`internal/services/url_service.go`)
The service layer handles the business logic:
- URL validation
- Short code generation
- Database operations
- URL redirection

#### Handlers (`internal/handlers/url_handler.go`)
The handler layer manages HTTP requests and responses:
- API endpoints
- Request validation
- Response formatting
- Error handling

## API Endpoints

### 1. Create Short URL
```http
POST /api/urls
Content-Type: application/json

{
    "url": "https://example.com/very/long/url"
}
```
Response:
```json
{
    "id": 1,
    "original_url": "https://example.com/very/long/url",
    "short_code": "AbCdEfGh",
    "created_at": "2024-03-15T10:30:00Z",
    "access_count": 0
}
```

### 2. Get All URLs
```http
GET /api/urls
```
Response:
```json
[
    {
        "id": 1,
        "original_url": "https://example.com/very/long/url",
        "short_code": "AbCdEfGh",
        "created_at": "2024-03-15T10:30:00Z",
        "access_count": 0
    }
]
```

### 3. Get URL by Short Code
```http
GET /api/urls/{shortCode}
```
Response:
```json
{
    "id": 1,
    "original_url": "https://example.com/very/long/url",
    "short_code": "AbCdEfGh",
    "created_at": "2024-03-15T10:30:00Z",
    "access_count": 0
}
```

## Implementation Details

### 1. URL Validation
The application validates URLs using Go's built-in `url.Parse` function:
```go
func ValidateURL(inputURL string) (string, error) {
    // Add https:// if no protocol is specified
    if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
        inputURL = "https://" + inputURL
    }
    
    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        return "", err
    }
    
    if parsedURL.Host == "" {
        return "", fmt.Errorf("invalid URL: missing host")
    }
    
    return parsedURL.String(), nil
}
```

### 2. Short Code Generation
Short codes are generated using a combination of UUID and base62 encoding:
```go
func generateShortCode() string {
    // Generate UUID
    id := uuid.New()
    
    // Convert to base62
    return base62.Encode(id[:])
}
```

### 3. Database Operations
The application uses GORM for database operations:
```go
// Initialize database
db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}

// Auto-migrate the schema
db.AutoMigrate(&models.URL{})
```

## Frontend Features

### 1. URL Shortening Form
- Input field for long URLs
- Submit button to create short URL
- Error handling and success messages

### 2. URL List
- Table displaying all shortened URLs
- Copy button for each short URL
- Access count tracking
- Created date display

## Usage Examples

### 1. Shortening a URL
1. Enter a long URL in the input field
2. Click "Shorten URL"
3. Copy the generated short code

### 2. Accessing a Shortened URL
1. Use the format: `http://localhost:8080/{shortCode}`
2. The application will redirect to the original URL
3. Access count will be incremented

### 3. Testing API Endpoints
1. Select an endpoint from the dropdown
2. Enter required data
3. Click "Test Endpoint"
4. View the response in the JSON viewer

## Error Handling

### 1. Invalid URLs
- Missing protocol (automatically adds https://)
- Invalid host
- Malformed URL structure

### 2. Database Errors
- Duplicate short codes
- Connection issues
- Transaction failures

### 3. API Errors
- Invalid request format
- Missing required fields
- Resource not found

## Security Considerations

### 1. URL Validation
- Protocol validation
- Host validation
- Length restrictions

### 2. Rate Limiting
- Request throttling
- IP-based restrictions
- Concurrent request handling

### 3. Data Protection
- SQL injection prevention
- XSS protection
- Input sanitization

## Performance Optimization

### 1. Caching
- In-memory caching for frequently accessed URLs
- Cache invalidation strategies
- Cache size management

### 2. Database Indexing
- Index on short_code for fast lookups
- Index on created_at for sorting
- Composite indexes for common queries

### 3. Response Optimization
- JSON response compression
- HTTP caching headers
- Connection pooling

## Deployment

### 1. Requirements
- Go 1.16 or higher
- SQLite3
- 512MB RAM minimum
- 1GB storage space

### 2. Environment Variables
```bash
GIN_MODE=release
PORT=8080
DB_PATH=urls.db
```

### 3. Build and Run
```bash
# Build the application
go build -o url-shortener cmd/main.go

# Run the application
./url-shortener
```

## Monitoring and Maintenance

### 1. Logging
- Request logging
- Error logging
- Access logging

### 2. Metrics
- URL creation rate
- Access patterns
- Error rates

### 3. Backup
- Database backup
- Configuration backup
- Log rotation

## Future Enhancements

### 1. Planned Features
- Custom short codes
- URL expiration
- Analytics dashboard
- User authentication

### 2. Performance Improvements
- Redis caching
- Load balancing
- CDN integration

### 3. Security Enhancements
- HTTPS enforcement
- API key authentication
- Rate limiting 