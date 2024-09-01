# The Meme Index API

The Meme Index API is a RESTful API for managing and retrieving memes. It is built using Go, Chi, and OpenAPI.

## Table of Contents

- [The Meme Index API](#the-meme-index-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Configuration](#configuration)
  - [API Documentation](#api-documentation)
  - [Development](#development)
    - [Prerequisites](#prerequisites)
    - [Running Locally](#running-locally)
    - [Code Generation](#code-generation)

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

- `PORT`: The port on which the server will run (default: `8080`).

You can set these variables in your environment or in a `.env` file.

## API Documentation

The API documentation is generated from the OpenAPI specification. You can view the Swagger UI by navigating to `/swagger/` in your browser once the server is running.

## Development

### Prerequisites

- Go 1.16 or later
- Taskfile

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