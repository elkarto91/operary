# Operary Phase 1 – Module Documentation

This document provides an extended reference for all Operary modules delivered in Phase 1. Each module exposes a set of REST APIs built with Go and MongoDB and is orchestrated by the core server in `backend/`.

## Final Operary+ Phase 1 Modules

| Module | Purpose | Status |
|-------|---------|-------|
| Core Operary | Tasks, shifts, audits – orchestrator core | ✅ Done |
| CorePad | Operator notes | ✅ Done |
| OpsMirror | Live dashboard/status monitoring | ✅ Done |
| AuditSync | Log operational audits | ✅ Done |
| PermitGrid | Manage permits and PTWs | ✅ Done |
| EquipTrust | Track asset/equipment lifecycle | ✅ Done |
| SensorVault | Log sensor alerts, anomalies | ✅ Done |
| TrainOps | Self-learning operational intelligence | ✅ Done |
| TraceBoard | RCA and fault tree tracking | ✅ Done |
| FlowGrid | Auto-assign shifts using skill matching | ✅ Done |
| CorePad++ | Multimedia notes & GPT-style escalations | ✅ Done |

---

## 1. Core Operary
The central API under `/v1` manages tasks and shifts. Handlers can be found in `backend/internal/handlers` and models under `backend/internal/models`.

- **Tasks** – CRUD endpoints exposed at `/v1/tasks`. Supports creation, listing, and status updates with audit logging.
- **Shifts** – Start and close shifts via `/v1/shifts`. Generates simple shift reports from task history.
- **Audit** – Retrieve audit logs at `/v1/audit`.
- **Webhook** – `/v1/webhook` auto‑creates tasks from external alerts.

## 2. CorePad
Structured note system for operators. Notes are stored in the `corepad_notes` collection.

- `POST /corepad/notes` – create a note with content, author and tags.
- `GET /corepad/notes/{id}` – retrieve a single note.

Metrics `corepad_notes_created_total` and `corepad_notes_failed_total` are exported for observability.

## 3. OpsMirror
Provides a real‑time dashboard view of system statuses.

- `GET /opsmirror/dashboard` – returns all monitored system states.
- Status records are stored in `opsmirror_status` and WARN counts are exposed via Prometheus under `opsmirror_status_warnings_total`.

## 4. AuditSync
Digital audit trail service used across modules.

- `POST /auditsync/entries` – submit an audit entry.
- `GET /auditsync/entries` – list all audits.

Audits reside in the `audit_entries` collection and are stamped with timestamps for traceability.

## 5. PermitGrid
Permit‑to‑work (PTW) request management.

- `POST /permitgrid/requests` – create a permit request.
- `GET /permitgrid/requests` – list all requests.
- `PATCH /permitgrid/requests/{id}/approve` – approve a request.
- `PATCH /permitgrid/requests/{id}/reject` – reject a request.

Each permit stores requester details, tags and approval metadata in the `permit_requests` collection.

## 6. EquipTrust
Tracks equipment check‑outs and returns.

- `POST /equiptrust/checkout` – record a check‑out entry.
- `PATCH /equiptrust/return/{id}` – mark equipment returned.
- `GET /equiptrust/ledger` – view the usage ledger.

Entries live in the `equipment_ledger` collection with timestamps for check‑out and return.

## 7. SensorVault
Stores sensor anomalies and telemetry events.

- `POST /sensorvault/events` – record a new sensor event.
- `GET /sensorvault/events` – list events.
- `GET /sensorvault/events/{id}` – retrieve a specific event.
- `DELETE /sensorvault/events/{id}` – delete an event (admin use).

Events are saved in `sensor_events` with optional links to tasks or audits.

## 8. TrainOps
Collects tagged operational notes for future model training.

- `POST /trainops/ingest` – add a training sample.
- `GET /trainops/predictions` – list stored samples.
- `POST /trainops/feedback` – submit feedback on a prediction.

## 9. TraceBoard
Root cause analysis and incident tracking.

- `POST /traceboard/incidents` – submit an incident report.
- `GET /traceboard/incidents` – list reports.
- `GET /traceboard/incidents/{id}` – fetch a single report.
- `DELETE /traceboard/incidents/{id}` – remove a report.

Reports in `traceboard_incidents` can later feed PDF exports and fault‑tree visualizations.

## 10. FlowGrid
Skill matching engine for shift scheduling.

- `POST /flowgrid/match-skills` – return best worker assignments based on required skills.
- `POST /flowgrid/schedule` – currently an alias of `match-skills`.

The module is stateless today but sets the foundation for automated scheduling.

## 11. CorePad++
An extension of CorePad to support multimedia attachments and GPT‑style escalation summaries. Implementation will build upon CorePad handlers and is scoped for future enhancement.

---

For configuration details and developer scripts see `operary_access_guide_updated.md`. The API specifications are available under `api_spec/` for import into Swagger or Postman.
