# Meme Index Web API

A Go-based web API for managing and searching memes, built using a **client → manager → accessor** architecture for clean separation of concerns.

---

## 🧩 Architecture Overview

### Layers

#### 1️⃣ Client (HTTP layer)
- Handles all incoming HTTP requests using **Fiber**.  
- Performs request parsing, basic validation, and marshals responses to JSON.
- Calls into the Manager layer for actual logic.
- Example folder: `internal/client/httpserver`

#### 2️⃣ Manager (Logic / Orchestration layer)
- Contains the application’s business logic.  
- Coordinates between multiple accessors (database, object storage, queues).  
- Enforces rules, normalization, and validation.
- Example folder: `internal/manager`

#### 3️⃣ Accessor (Integration layer)
- Encapsulates connections to backend services:
  - **Database** (Postgres via `pgx`): stores memes, tags, file metadata.
  - **Queue** (Redis): background jobs (e.g., thumbnailing, AI tagging).
  - **Object Store** (S3/MinIO): stores the actual meme files.
- Example folder: `internal/accessor`

Each layer only talks **downwards** — the Client never directly touches accessors.

```
client (Fiber HTTP)
   ↓
manager (logic, validation)
   ↓
accessor (Postgres, Redis, S3)
```

---

## 📁 Core Concepts

### Memes
Every uploaded meme has:
- A unique `id`
- Metadata (filename, size, mime type, created_at, etc.)
- One or more **tags**

### Tags
Tags are stored centrally in a `tags` table:
- `name` (unique, lowercase)
- `kind` (`manual`, `ai`, `system`)
- `created_at`

Memes and tags are connected through `file_tags` (many-to-many).

### Search
Search can filter memes by:
- Tags (`tags_all`, `tags_any`, `tags_not`)
- Future: text query, mime type, date range, or visual similarity.

Tag suggestions while typing (`GET /v1/tags/suggest?q=fun`) use Postgres trigram indexes for fast fuzzy/prefix matching.

---

## 🚀 Current Roadmap

| Phase | Goal | Focus |
|-------|------|--------|
| 1 | **Core API Skeleton** | Layers, configuration, health endpoint |
| 2 | **Tags & Search** | Tag model, search/filtering, suggestions |
| 3 | **Uploads & Presign** | S3/MinIO upload flow |
| 4 | **Background Jobs** | Redis + Worker pipeline |
| 5 | **AI Tagging (optional)** | Auto-generated tags, NSFW filtering, etc. |

---

## 🧠 Tech Stack

- **Language:** Go 1.22+
- **Framework:** [Fiber](https://gofiber.io/)
- **Database:** PostgreSQL (metadata, tags)
- **Queue:** Redis
- **Object Storage:** MinIO / S3 compatible
- **Docs:** OpenAPI 3 (Swagger / Redoc)

---

## 🛠️ Development Setup

```bash
go mod tidy
mkdir -p cmd/api internal/{client,httpserver,manager,accessor,models,config,errs} docs
go run ./cmd/api
```

Environment variables (sample):
```bash
DATABASE_URL=postgres://app:app@localhost:5432/memeindex?sslmode=disable
REDIS_ADDR=localhost:6379
S3_ENDPOINT=http://localhost:9000
S3_BUCKET=memeindex
S3_ACCESS_KEY=minio
S3_SECRET_KEY=minio12345
S3_USE_PATH_STYLE=true
```

---

## 📄 License
MIT
