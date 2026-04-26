# Basic Go Web Service

A simple HTTP web service written in Go with multiple endpoints for demonstration purposes.

## Features

- **Hello Endpoint**: Returns a greeting message
- **Health Check Endpoint**: Verifies service is running
- **Echo Endpoint**: Echoes back messages sent by clients
- JSON request/response handling
- Clean error handling and HTTP status codes

## Prerequisites

- Go 1.21 or higher installed
- Basic knowledge of Go and HTTP

## Getting Started

### Installation

1. Clone the repository:
```bash
git clone https://github.com/AleBoi/newgitproj.git
cd newgitproj
```

2. Download dependencies:
```bash
go mod download
```

### Running the Service

Start the web service:
```bash
go run main.go
```

The service will start on `http://localhost:8080`

You should see:
```
Starting web service on http://localhost:8080
Available endpoints:
  GET  /hello
  GET  /health
  POST /echo
```

## API Endpoints

### 1. Hello Endpoint
**Request:**
```bash
curl http://localhost:8080/hello
```

**Response:**
```json
{
  "status": "success",
  "message": "Hello from Go web service!",
  "data": {
    "version": "1.0",
    "service": "Basic Go Web Service"
  }
}
```

### 2. Health Check Endpoint
**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "success",
  "message": "Service is healthy",
  "data": {
    "status": "healthy",
    "uptime": "running"
  }
}
```

### 3. Echo Endpoint
**Request:**
```bash
curl -X POST http://localhost:8080/echo \
  -H "Content-Type: application/json" \
  -d '{"message":"Hello World"}'
```

**Response:**
```json
{
  "status": "success",
  "message": "Echo response",
  "data": {
    "echo": "Hello World"
  }
}
```

## Building for Production

Create an executable binary:
```bash
go build -o web-service
```

Run the binary:
```bash
./web-service
```

## Project Structure

```
newgitproj/
├── main.go          # Main application file with all handlers
├── go.mod           # Go module definition
└── README.md        # This file
```

## Error Handling

The service handles common errors:
- Invalid HTTP methods return `405 Method Not Allowed`
- Invalid JSON payloads return `400 Bad Request`
- All responses include a status field for easy parsing

## License

MIT

## Author

AleBoi
