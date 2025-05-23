# 📘 Operary – Product Requirements Document (PRD)

> “Product is the act of structuring chaos into trustable, usable loops.”

---

## 🔧 Product Name: **Operary**

A coordination engine for frontline operations. It bridges human workflows, machine signals, and shift-based action logs.

---

## 🎯 Objective

Design and build a **modular, API-first SaaS product** that:

- Orchestrates real-time tasks between humans and systems  
- Tracks incident-response workflows across shifts  
- Provides a traceable, auditable log of all operational execution  
- Integrates with SCADA/MES/ERP for signal and status input/output  
- Offers a minimal, responsive frontend (optional) for mobile/terminal access

---

## 👥 Personas

| Persona | Role |
|--------|------|
| 👷 Operator | Receives task instructions, closes or flags incidents  
| 🧭 Supervisor | Assigns tasks, resolves escalations, documents shift  
| 🛠️ Maintenance Engineer | Handles escalated issues, leaves structured logs  
| 📋 Compliance Lead | Reviews shift logs and incident trails for audits  
| 🧠 Admin / System Integrator | Manages Org, API keys, integrates with other systems

---

## 📦 Features & Requirements

### 1. Task Orchestration Engine (MVP)

| Feature | Description |
|--------|-------------|
| Create Task | Manual or machine-triggered  
| Assign Task | Based on roles and availability  
| Complete Task | Close or escalate with notes, evidence  
| Shift Logging | Group tasks + comments into shift report  
| API Access | REST API to create/read/update task states  
| Roles & Org | Basic Org + role-based access for users

---

### 2. Audit & Escalation Trails

- Every task change must include: who, when, reason, linked upstream  
- Escalation logs = traceable up to incident root  
- All actions are **read-only** for audit once completed

---

### 3. Event Ingestion + External Triggers

| Source | Integration |
|--------|-------------|
| SCADA or Sensors | Webhook/REST-based alerts into Operary  
| MES / ERP | Optional polling or push  
| Manual Entry | Supervisor input during handoff or alert

---

## 🧰 Technical Requirements

| Component | Spec |
|-----------|------|
| API Design | REST-first (OpenAPI 3.0)  
| Backend Language | Golang  
| Storage | PostgreSQL or SQLite for demo  
| Auth | Token-based, scoped per Org  
| Hosting | Docker + Local run + GitHub Pages for docs

---

## 🧪 KPIs (v0.1 Demo)

| Metric | Target |
|--------|--------|
| Task creation → close loop time | < 2 minutes  
| API test coverage | 90%  
| Audit trail completeness | 100% trace from event to resolution  
| Documentation | Swagger + Markdown clarity

---

## 🚦 User Flows (MVP)

1. Machine sends alert → webhook triggers task
2. Task routes to on-duty role (e.g. “Line Lead”)
3. Lead adds notes → closes or escalates
4. Maint Eng logs final action
5. Shift ends → report auto-generated and sealed
6. Audit Lead opens report in read-only archive

---

## 🧱 Out of Scope (for v0.1)

- Role scheduling / calendar sync  
- Full UI/UX dashboard  
- Analytics & predictive AI  
- Multilingual support  
- Voice inputs / OCR / hardware API

---

## ⚠️ Known Risks

- False triggers or alert spam  
- Token permissions abuse  
- Audit lockout before data sync  
- Shift log overwrite unless sealed  
- SCADA integrations vary by plant

---

## 🧩 Milestones

| Milestone | Output |
|-----------|--------|
| ✅ Product Management | Ideation, Business Case, PRD  
| 🧱 Tech Spec | architecture.md + OpenAPI YAML  
| 💻 Backend | core Golang REST endpoints  
| 🔐 Auth & Org Logic | basic API token guard  
| 📄 Demo Docs | test script, curl examples  
| 🧪 Postman Collection | for consumers and validators

---

## 📁 Related

- [01-ideation.md](./01-ideation.md)  
- [02-business-case.md](./02-business-case.md)  
- [docs/architecture.md](../docs/architecture.md)  
- [api-spec/openapi.yaml](../api-spec/openapi.yaml)

---

> This document is the first interface of belief. It defines what we’re building — and why.

