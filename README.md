# LBS Service

LBS Service is a unified API service that integrates multiple map services, including Baidu Map, Gaode Map, and Tencent Map. This service is built using Go and the Gin framework, with Redis as a caching middleware to enhance performance.

## Features

- Unified API for accessing different map services
- Automatic switching between Baidu, Gaode, and Tencent maps
- Caching responses using Redis to improve response times
- Easy to extend and maintain

## Project Structure

```
lbs-service
├── src
│   ├── controllers          # Contains API request handlers
│   ├── middlewares          # Contains middleware for Redis caching
│   ├── services             # Contains services for different map APIs
│   ├── utils                # Contains utility functions for map data processing
│   ├── main.go              # Entry point of the application
│   └── config               # Contains configuration settings
├── go.mod                   # Module definition and dependencies
├── go.sum                   # Checksums for module dependencies
└── README.md                # Project documentation
```

## Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd lbs-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Configure your API keys and Redis connection details in `src/config/config.go`.

## Usage

1. Start the server:
   ```
   go run src/main.go
   ```

2. Access the API endpoints:
   - Example endpoint to get map data:
     ```
     GET /api/map/data?location=<location>
     ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.