# Operary – Module & Function Overview

This document summarizes the backend modules found in the `backend/internal` directory and the key functions provided by each package. It is intended as a high-level reference when exploring the codebase.

## CorePad
Operator note system for capturing shift observations.

**Main handlers**
- `CreateNoteHandler` – POST `/corepad/notes`
- `GetNoteHandler` – GET `/corepad/notes/{id}`

**Usecase layer**
- `CreateNote(content, author, tags)`
- `GetNote(id)`

**Repository layer**
- `SaveNote(Note)` and `GetNoteByID(id)`

## OpsMirror
Dashboard showing current state of critical systems.

**Main handler**
- `GetStatusDashboard` – GET `/opsmirror/dashboard`

**Usecase layer**
- `FetchStatuses()`

**Repository layer**
- `GetAllStatuses()`
- `UpsertStatus(status)`

## AuditSync
Digital audit trail service.

**Main handlers**
- `CreateAuditEntry` – POST `/auditsync/entries`
- `GetAuditEntries` – GET `/auditsync/entries`

**Usecase layer**
- `CreateAudit(request)`
- `ListAudits()`

**Repository layer**
- `SaveAudit(AuditEntry)`
- `ListAudits()`

## PermitGrid
Permit-to-work request management.

**Main handlers**
- `CreatePermitRequest` – POST `/permitgrid/requests`
- `ListPermitRequests` – GET `/permitgrid/requests`
- `ApprovePermit` – PATCH `/permitgrid/requests/{id}/approve`
- `RejectPermit` – PATCH `/permitgrid/requests/{id}/reject`

**Usecase layer**
- `CreatePermit(request)`
- `ListPermits()`
- `ApprovePermit(id, approver)`
- `RejectPermit(id, approver)`

**Repository layer**
- `SavePermit(Permit)`
- `ListPermits()`
- `UpdatePermitStatus(id, status, approver)`

## EquipTrust
Equipment handover ledger.

**Main handlers**
- `CheckoutEquipment` – POST `/equiptrust/checkout`
- `ReturnEquipment` – PATCH `/equiptrust/return/{id}`
- `ListLedger` – GET `/equiptrust/ledger`

**Usecase layer**
- `Checkout(request)`
- `Return(id, time)`
- `ListLedger()`

**Repository layer**
- `Save(EquipmentLedger)`
- `UpdateReturn(id, time)`
- `List()`

## SensorVault
Recording machine events and telemetry.

**Main handlers**
- `RecordEvent` – POST `/sensorvault/events`
- `ListEvents` – GET `/sensorvault/events`

**Usecase layer**
- `RecordEvent(request)`
- `GetEvents()`

**Repository layer**
- `Save(SensorEvent)`
- `List()`

## TrainOps
Self-learning engine that stores tagged notes and feedback.

**Main handlers**
- `IngestData` – POST `/trainops/ingest`
- `ListPredictions` – GET `/trainops/predictions`
- `ProvideFeedback` – POST `/trainops/feedback`

**Usecase layer**
- `Ingest(request)`
- `List()`
- `AddFeedback(id, fb)`

**Repository layer**
- `Save(TrainingSample)`
- `List()`
- `UpdateFeedback(id, fb)`

## Shared Models
Core task and shift logic lives under `internal/models`:
- `Task` with functions `CreateTask`, `GetAllTasks`, `UpdateTask`, `GetTasksByTimeRange`.
- `Shift` with functions `StartShift`, `CloseShift`, `GetShiftByID`.
- `AuditLog` utilities such as `RecordAudit` and `GetAuditLogs`.
- `GenerateShiftReport` for simple shift reports.

## Services
Background workers live in `internal/services`.
- `StartNotificationService(logger)` prints periodic heartbeats and represents future integrations.

For endpoint details see the OpenAPI specifications in the `api_spec/` directory.
