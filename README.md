# Golang Starter Project

This is a starter project for a Golang application with a focus on a clean architecture, hot-reloading for development, and easy setup with Docker.

## Technologies Used

*   **Backend:** Golang
*   **Database:** PostgreSQL
*   **Hot-Reloading:** [air](https://github.com/air-verse/air)
*   **Containerization:** Docker and Docker Compose

## Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.18 or higher)
*   [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Without Docker

1.  **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd golang-starter
    ```

2.  **Install dependencies:**

    ```bash
    go mod download
    ```

3.  **Install `air` for hot-reloading:**

    ```bash
    go install github.com/air-verse/air@latest
    ```

4.  **Set up environment variables:**

    Create a `.env` file in the root of the project and add the following variables:

    ```
    DATABASE_URL=postgresql://user:password@localhost:5432/mydatabase?sslmode=disable
    ENV=development
    ```

5.  **Run the application:**

    ```bash
    air
    ```

    The application will be running at `http://localhost:8080`.

### With Docker

1.  **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd golang-starter
    ```

2.  **Build and run the application with Docker Compose:**

    ```bash
    docker-compose up --build
    ```

    The application will be running at `http://localhost:8080`. The database will be running on port `5432`, and pgAdmin will be accessible at `http://localhost:5050`.

## Hot Reloading

This project is configured for hot-reloading using `air`. When you make changes to any `.go` files, the application will automatically restart.

*   **Local Development:** When running the application with `air` locally, hot-reloading is enabled by default.
*   **Docker Development:** When running the application with `docker-compose`, hot-reloading is also enabled. The `volumes` in the `docker-compose.yml` file mount the local code into the container, and `air` watches for changes.

## Project Structure

```
.
├── cmd
│   └── main.go         # Entry point of the application
├── configs             # Configuration files
├── internal
│   ├── controllers     # HTTP handlers
│   ├── middlewares     # Middlewares for the router
│   ├── models          # Database models
│   ├── repositories    # Database queries
│   ├── routes          # Application routes
│   └── services        # Business logic
├── pkg
│   └── database        # Database connection
├── .air.toml           # Configuration for air (hot-reloading)
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile          # Dockerfile for the application
└── go.mod              # Go modules
```
