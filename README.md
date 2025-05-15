# URL Shortener Service

A modern, efficient URL shortening service built with Go and Gin framework. This service provides a RESTful API and a clean web interface for shortening URLs, tracking click statistics, and managing shortened URLs.

## Features

- 🔗 URL shortening with custom short codes
- 📊 Click tracking and statistics
- 🎯 Clean and responsive web interface
- 🔒 Input validation and error handling
- 📱 Mobile-friendly design
- 🚀 Fast and efficient performance
- 💾 SQLite database for data persistence

## Tech Stack

- **Backend**: Go (Golang)
- **Web Framework**: Gin
- **Database**: SQLite with GORM
- **Frontend**: HTML, CSS (Bootstrap 5), JavaScript
- **Package Management**: Go Modules

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener-go.git
   cd url-shortener-go
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:5001`

## Project Structure

```
url-shortener-go/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── handlers/         # HTTP request handlers
│   ├── models/          # Data models
│   └── services/        # Business logic
├── web/
│   └── templates/       # HTML templates
└── go.mod              # Go module file
```

## API Endpoints

### Web Interface
- `GET /` - Web interface for URL shortening

### API Endpoints
- `GET /api/urls` - Get all shortened URLs
- `POST /api/urls` - Create a new shortened URL
- `GET /api/urls/:shortCode` - Get a specific shortened URL
- `PUT /api/urls/:shortCode` - Update a shortened URL
- `DELETE /api/urls/:shortCode` - Delete a shortened URL
- `GET /api/urls/:shortCode/stats` - Get statistics for a shortened URL
- `GET /:shortCode` - Redirect to the original URL

## Usage

1. Open your web browser and navigate to `http://localhost:5001`
2. Enter a URL in the input field
3. Click "Shorten" to generate a shortened URL
4. Use the shortened URL to redirect to the original URL
5. View click statistics and manage your shortened URLs

## Features in Detail

### URL Shortening
- Automatically adds `https://` if no protocol is specified
- Validates URLs before shortening
- Generates unique short codes using UUID

### Click Tracking
- Tracks the number of clicks for each shortened URL
- Provides click statistics through the API
- Updates click count in real-time

### Web Interface
- Clean and intuitive design
- Responsive layout for all devices
- Real-time updates
- Collapsible URL list
- Error handling and success notifications

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Bootstrap](https://getbootstrap.com/) 