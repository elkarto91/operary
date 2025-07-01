# Operary Platform – Access & Configuration Guide

This document provides a complete overview of how to access the Operary platform services locally and how they are configured via Docker and environment variables.

---

## 📍 1. Platform Endpoints

| Feature             | Endpoint                                     | Description                                  |
|---------------------|----------------------------------------------|----------------------------------------------|
| Core Operary API    | `/shifts`, `/tasks`                          | Main Operary orchestration APIs              |
| CorePad API         | `/corepad/notes`                             | Operator note-taking system                  |
| OpsMirror API       | `/opsmirror/dashboard`                       | Real-time system status dashboard            |
| AuditSync API       | `/auditsync/entries`                         | Digital audits and findings trail            |
| PermitGrid API      | `/permitgrid/requests`                       | PTW (Permit to Work) request system          |
| EquipTrust API      | `/equiptrust/checkout`, `/return/{id}`      | Equipment handover ledger                    |
| Prometheus Metrics  | `/metrics`                                   | Monitoring data for all modules              |
| MongoDB (Container) | `mongo:27017`                                | Internal MongoDB for services                |
| MongoDB (Host)      | `localhost:27017`                            | Accessible via Compass or CLI                |

---

## 🔐 2. Configuration and Credentials

### Docker Compose (`backend/docker-compose.yml`)

```yaml
environment:
  - MONGO_URI=mongodb://mongo:27017
  - MONGO_DB=operary_dev
  - ORG_TOKEN=demo-org-token
```

| Variable Name   | Value                        | Description                      |
|------------------|-----------------------------|----------------------------------|
| `MONGO_URI`      | `mongodb://mongo:27017`     | MongoDB connection string        |
| `MONGO_DB`       | `operary_dev`               | Database used by all modules     |
| `ORG_TOKEN`      | `demo-org-token`            | (Optional) Demo organization ID  |

> ⚠️ MongoDB currently has **no authentication** for local dev. Use credentials for production.

---

## 🧪 3. API Specs (Swagger / Postman)

Import these OpenAPI specs into [Swagger Editor](https://editor.swagger.io/) or Postman:

- `api_spec/openapi.yaml` (Core Operary)
- `api_spec/corepad.yaml`
- `api_spec/opsmirror.yaml`
- `api_spec/auditsync.yaml`
- `api_spec/permitgrid.yaml`
- `api_spec/equiptrust.yaml`

---

## 📊 4. Prometheus Metrics Overview

| Metric Name                           | Module       | Description                              |
|---------------------------------------|--------------|------------------------------------------|
| `corepad_notes_created_total`         | CorePad      | Number of notes created                  |
| `corepad_notes_failed_total`          | CorePad      | Failed note submissions                  |
| `opsmirror_status_warnings_total`     | OpsMirror    | WARN states visible on dashboard         |
| `auditsync_total_submitted`           | AuditSync    | Total audits submitted                   |
| `auditsync_submission_failures_total` | AuditSync    | Failed audit submissions                 |
| `permitgrid_total_submitted`          | PermitGrid   | Permits requested                        |
| `permitgrid_total_approved`           | PermitGrid   | Permits approved                         |
| `permitgrid_total_rejected`           | PermitGrid   | Permits rejected                         |
| `equiptrust_checkouts_total`          | EquipTrust   | Equipment checked out                    |
| `equiptrust_returns_total`            | EquipTrust   | Equipment returned                       |
| `operary_total_requests`              | Global       | All HTTP requests                        |
| `operary_uptime_seconds`              | Global       | Server uptime                            |
| `operary_shifts_started`              | Global       | Shifts started                 |
| `operary_shifts_closed`               | Global       | Shifts closed                  |
| `operary_tasks_created`               | Global       | Tasks created via API          |

---

## 🛠 5. Developer Scripts & Tools

| Script                        | Purpose                                      |
|-------------------------------|----------------------------------------------|
| `simulate_opsmirror_status.go`| Simulate real-time OpsMirror alerts          |
| `seed_auditsync.go`           | Insert fake audits into AuditSync            |
| `seed_permitgrid.go`          | Insert demo permits into PermitGrid          |
| `seed_equiptrust.go`          | Insert dummy equipment usage into ledger     |

---

## ✅ Ready to Use

Start services with:

```bash
cd backend
docker-compose up --build
```

Test endpoints, metrics, and UI via Postman, Swagger, and Prometheus.

---