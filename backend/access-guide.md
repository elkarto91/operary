# Operary Platform – Access & Configuration Guide

This document provides a complete overview of how to access the Operary platform services locally and how they are configured via Docker and environment variables.

---

## 📍 1. Platform Endpoints

| Feature             | Endpoint                                     | Description                                  |
|---------------------|----------------------------------------------|----------------------------------------------|
| Core Operary API    | `http://localhost:8080/shifts`, `/tasks`     | Main Operary orchestration APIs              |
| CorePad API         | `http://localhost:8080/corepad/notes`        | Create and fetch operator notes              |
| OpsMirror API       | `http://localhost:8080/opsmirror/dashboard`  | View live system statuses                    |
| Prometheus Metrics  | `http://localhost:8080/metrics`              | Exposes uptime, usage, and alert counters    |
| MongoDB (Container) | `mongo:27017`                                | For services inside Docker                   |
| MongoDB (Host)      | `localhost:27017`                            | For local tools like Compass or mongosh      |

---

## 🔐 2. Configuration and Credentials

### 🔧 Docker Compose (`backend/docker-compose.yml`)

```yaml
services:
  mongo:
    image: mongo:6
    container_name: operary-mongo
    ports:
      - "27017:27017"
    volumes:
      - mongo-data:/data/db

  api:
    build:
      context: .
    container_name: operary-api
    ports:
      - "8080:8080"
    depends_on:
      - mongo
    environment:
      - MONGO_URI=mongodb://mongo:27017
      - MONGO_DB=operary_dev
      - ORG_TOKEN=demo-org-token

volumes:
  mongo-data:
```

### 📌 Active Environment Variables

| Variable Name       | Value                        | Purpose                              |
|---------------------|------------------------------|--------------------------------------|
| `MONGO_URI`         | `mongodb://mongo:27017`      | MongoDB connection URI               |
| `MONGO_DB`          | `operary_dev`                | Name of the MongoDB database         |
| `ORG_TOKEN`         | `demo-org-token`             | Demo token for organizational scope  |

> ⚠️ **Authentication:** MongoDB is currently open (no auth). To secure for production, add:
```yaml
environment:
  - MONGO_INITDB_ROOT_USERNAME=admin
  - MONGO_INITDB_ROOT_PASSWORD=securepass
```
Then update:
```
MONGO_URI=mongodb://admin:securepass@mongo:27017
```

---

## 🧪 3. API Specs (Swagger / Postman)

You can import these files into Swagger Editor or Postman for API exploration:

- `/api_spec/corepad.yaml`
- `/api_spec/opsmirror.yaml`
- `/api_spec/openapi.yaml` (core Operary APIs)

---

## 📊 4. Prometheus Metrics Available

| Metric                                | Description                              |
|---------------------------------------|------------------------------------------|
| `corepad_notes_created_total`         | Number of notes created                  |
| `corepad_notes_failed_total`          | Number of failed note creations          |
| `opsmirror_status_warnings_total`     | Systems currently in WARN state          |
| `operary_total_requests`              | Total HTTP requests received             |
| `operary_uptime_seconds`              | Application uptime in seconds            |

---

## 🛠 5. Developer Tools & Scripts

| Script                                | Description                              |
|---------------------------------------|------------------------------------------|
| `scripts/simulate_opsmirror_status.go`| Simulates live system statuses for OpsMirror |
| `scripts/sim_shift_test.go`           | Simulates a shift pattern (optional/dev) |

---

## ✅ Ready to Use

Start with:
```bash
cd backend
docker-compose up --build
```

Test Core Operary, CorePad, OpsMirror, metrics, and database integration locally.

---