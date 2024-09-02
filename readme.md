# The Meme Index API

The Meme Index API is a RESTful API for managing and retrieving memes. It is built using Go, Chi, and OpenAPI.

## Table of Contents

- [The Meme Index API](#the-meme-index-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Configuration](#configuration)
    - [Server Configuration](#server-configuration)
    - [PostgreSQL Configuration](#postgresql-configuration)
    - [MinIO Configuration](#minio-configuration)
  - [API Documentation](#api-documentation)
  - [Development](#development)
    - [Prerequisites](#prerequisites)
    - [Running Locally](#running-locally)
    - [Code Generation](#code-generation)
  - [TODO's](#todos)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/username/the-meme-index-api.git
    cd the-meme-index-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

1. Generate the API code from the OpenAPI specification:

    ```bash
    task apigen
    ```

2. Run the API server:

    ```bash
    task run
    ```

## Configuration

The API can be configured using environment variables. The following variables are available:

### Server Configuration
- `PORT`: The port on which the server will run (default: `8080`).

### PostgreSQL Configuration
- `PG_DB_HOST`: The hostname of the PostgreSQL database (e.g., `localhost`).
- `PG_DB_PORT`: The port number of the PostgreSQL database (e.g., `5432`).
- `PG_DB_USER`: The username for the PostgreSQL database (e.g., `postgres`).
- `PG_DB_PASS`: The password for the PostgreSQL database (e.g., `password`).
- `PG_DB_NAME`: The name of the PostgreSQL database (e.g., `memeindex`).
- `PG_DB_SSL_MODE`: The SSL mode for the PostgreSQL database (e.g., `disable`).
### MinIO Configuration
- `MINIO_ENDPOINT`: The endpoint for the MinIO server (e.g., `localhost:9000`).
- `MINIO_ACCESS_KEY`: The access key for the MinIO server (e.g., `minio`).
- `MINIO_ACCESS_KEY_ID`: The access key ID for the MinIO server (e.g., `minio123`).
- `MINIO_BUCKET_NAME`: The bucket name in the MinIO server (e.g., `memeindex`).
- `MINIO_USE_SSL`: Whether to use SSL with the MinIO server (e.g., `false`).

You can set these variables in your environment or in a .env file.

## API Documentation

The API documentation is generated from the OpenAPI specification. You can view the Swagger UI by navigating to `/swagger/` in your browser once the server is running.

## Development

### Prerequisites

- Go 1.16 or later
- Taskfile
- PostgresSQL instance
- Minio Instance

### Running Locally

1. Install Taskfile:

    ```bash
    go get -u github.com/go-task/task/v3/cmd/task
    ```

2. Run the server:

    ```bash
    task run
    ```

### Code Generation

To generate the API code from the OpenAPI specification, run:

```bash
task apigen
```

## TODO's

- [ ] Add Logging for failed requests
- [ ] Create Docker files for Minio and Postgres
- [ ] Add Tests
  - [ ] Accessor tests
  - [ ] Manager tests
  - [ ] Client tests
- [ ] TBD - (I will probably come up with more things later)