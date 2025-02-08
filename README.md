# Chrome Extension API

This project provides a backend API for a Chrome extension. The API is built using Go and provides endpoints to handle requests from the Chrome extension.

## Endpoints

### POST /api/v1/requests

This endpoint handles incoming requests from the Chrome extension.

**Request Example:**

```sh
curl -X POST http://localhost:8080/api/v1/requests -H "Content-Type: application/json" -d '{"link": "https://google.com"}'
```

**Response Example:**

```json
{
  "html": ""
}
```

### OPTIONS /api/v1/requests

This endpoint handles preflight requests for CORS.

**Request Example:**

```sh
curl -X OPTIONS http://localhost:8080/api/v1/requests
```

**Response Example:**

```sh
HTTP/1.1 200 OK
Allow: OPTIONS, POST
```

## Running with Docker Compose

To run the API using Docker Compose, follow these steps:

1. **Build and run the Docker container:**

```sh
docker compose up --build
```

2. **Access the API:**

The API will be available at `http://localhost:8080`.

## Project Structure

- `app/`: Contains the application code.
- `app/controllers/`: Contains the controllers for handling API requests.
- `app/controllers/api/v1/`: Contains the version 1 API controllers.
- `main.go`: The entry point of the application.
- `go.mod`: The Go module file.
- `go.sum`: The Go dependencies file.
- `Dockerfile`: The Dockerfile for building the Docker image.
- `docker compose.yml`: The Docker Compose file for managing the Docker container.

## Environment Variables

- `PORT`: The port on which the HTTP server will run. Default is `8080`.

## License

This project is licensed under the MIT License.
