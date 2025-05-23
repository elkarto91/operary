# 🧱 Operary – Technical Architecture

> A REST-first coordination engine for operational tasks, shift flows, and traceable action logs.

---

## 🔧 Stack Overview (MVP v0.1)

| Layer | Component |
|-------|-----------|
| 🧠 Backend | **GoLang** (fiber or chi router)  
| 🗃️ Database | **MongoDB** (JSON-like task, audit, shift docs)  
| 🔐 Auth | Bearer tokens (Org-scoped)  
| 📡 API Layer | RESTful endpoints + optional Webhooks  
| 📦 Event Ingestion | POST hooks + cron-like trigger ingestion  
| 📝 Storage | MongoDB task_logs, shift_logs, audit_docs  
| 🖥️ Admin Panel (optional) | Static frontend or Swagger UI  
| 🔁 CI/CD | GitHub Actions + Docker build  
| 📓 Docs | Swagger (OpenAPI) + Markdown specs

---

## 📁 Proposed MongoDB Collections

| Collection | Fields |
|------------|--------|
| `tasks` | id, title, source, assigned_to, status, due_time, escalation_id, notes[], created_by, created_at  
| `users` | id, name, role, org_id, is_active  
| `shifts` | id, started_at, ended_at, supervisor_id, closed, tasks[], notes  
| `audit_logs` | entity_type, entity_id, action, user_id, timestamp, delta_snapshot  
| `orgs` | id, name, contact, api_token, integrations[]

---

## 🧠 Backend Architecture (High Level)

```txt
+---------------------+       +----------------+
| External Systems    | --->  |  Webhook/API   |
| (SCADA/MES/Manual)  |       |  Ingestion     |
+---------------------+       +--------+-------+
                                     |
                              +------+------+
                              |   API Router |
                              | (REST/JSON)  |
                              +------+------+
                                     |
              +----------------------+------------------+
              |                                          |
        +-----v-----+                             +------v-------+
        | Task      |                             |  Shift       |
        | Engine    |                             |  Log Builder |
        +-----------+                             +--------------+
              |                                          |
        +-----v---------+                        +-------v---------+
        | MongoDB Layer |                        | Audit Service   |
        +---------------+                        +-----------------+

# 🔐 API Auth Strategy

* Each Org has a scoped `api_token`
* Token grants access to:
    * Task creation
    * Status updates
    * Audit reads (if sealed)
* Org admin can rotate token
* Role-based user creation (e.g., operator, supervisor, auditor)

---

# 🧪 Observability (MVP)

| Method      | Tool       | Logs                           | Monitoring             | Audit                        | Testing                               |
| :---------- | :--------- | :----------------------------- | :--------------------- | :--------------------------- | :------------------------------------ |
| **Logs** | Structured | JSON logs per task event       | Local logs + API timing | Stored in `audit_logs` collection | Postman collection + curl test scripts |

---

# ⚠️ Known Challenges

| Area                    | Risk                                             |
| :---------------------- | :----------------------------------------------- |
| **Audit Immutability** | Mongo is mutable by default. Must enforce locks at app level. |
| **External Webhooks** | Debounce + idempotency checks required           |
| **Shift Closure Race Conditions** | Supervisor may seal shift before tasks complete  |
| **Schema Drift** | Use central validation or typed models           |

---

# 🔄 Alternatives Considered

| Option     | Reason for rejection                          |
| :--------- | :-------------------------------------------- |
| **PostgreSQL** | Rigid schema; event log modeling harder than document store |
| **Firebase** | Not self-hostable for industrial enterprise POCs |
| **gRPC** | Adds complexity vs plain REST for MVP         |
| **Full UI-first** | Increases dev scope unnecessarily for backend demo |