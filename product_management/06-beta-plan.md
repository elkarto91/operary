# 🧪 Operary – Beta Plan

> “A great beta doesn’t just test functionality — it tests belief.”

---

## 🧭 Beta Objective

Validate that Operary’s core task coordination, shift tracking, and audit logging:

- Reduces ambiguity during shifts  
- Produces a usable audit trail automatically  
- Integrates with signal-triggered tasks  
- Helps supervisors feel more in control, not less

---

## 🎯 Target Beta Scope

| Target Group | Description |
|--------------|-------------|
| 🔹 Internal Team | Simulated test with systemsignal.dev contributors (backend, PM)  
| 🔸 Industrial Partner (Simulated) | Mocked site: energy ops, factory floor, utility grid team  
| 🛠️ Real Pilot (stretch) | Partner org or lab familiar with SCADA or shift ops

---

## 🧑‍💼 Personas in Beta

| Role | Tasks |
|------|-------|
| 👷 Operator | Receives + completes tasks  
| 🧭 Supervisor | Starts/ends shifts, assigns, escalates  
| 🛠️ Maintenance | Handles escalations, adds resolution notes  
| 📋 Auditor | Queries audit logs and views sealed shift reports

---

## 📦 What Will Be Tested

| Feature | Test Objective |
|---------|----------------|
| Task creation/assignment | Is it fast, intuitive, traceable?  
| Shift start/close | Is the report generated? Is the handoff clear?  
| Audit trail | Can every action be explained and proven?  
| External signal | Can we simulate SCADA → task mapping correctly?  
| Auth + Org | Are tokens scoped correctly?  
| Role behavior | Can personas access only their intended actions?

---

## 🧪 Success Metrics

| Metric | Target |
|--------|--------|
| ⏱️ Task creation to closure time | < 5 mins end-to-end  
| 📄 Audit completeness | 100% traceability on task flows  
| 🧾 Shift report generation | PDF accuracy + timestamp alignment  
| 📈 System uptime during test window | 95%+  
| 📬 Feedback from 3+ test users | Form + structured debrief

---

## 📑 Beta Assets

| Asset | Location |
|-------|----------|
| Test user guide | `/docs/beta-user-guide.md`  
| Postman collection | `/tests/postman_collection.json`  
| Shift simulation script | `/scripts/sim_shift_test.go`  
| Beta feedback form | Notion/Markdown format  
| Audit report samples | `/tests/samples/reports/`  

---

## 📆 Timeline

| Week | Activities |
|------|------------|
| Week 1 | Dry run: internal tasks + shift logs  
| Week 2 | Add mock signal trigger → webhook test  
| Week 3 | Invite feedback from testers  
| Week 4 | Clean audit logs, export test reports  
| Week 5 | Retrospective and refine roadmap

---

## 🛠️ Dev/Infra Requirements

- MongoDB instance  
- Hosted Go backend (localhost or fly.io)  
- Exposed endpoint (via ngrok, fly.io, or Vercel proxy)  
- Postman or curl access  
- Optional: Grafana or external log viewer

---

## ✅ Exit Criteria for Beta

| Goal | Threshold |
|------|-----------|
| Real use case tested | 2+ tasks, 1+ shift log, 1+ escalation  
| No major data loss or crash | System handled >95% of flows  
| Feedback received | At least 2 structured inputs  
| Postmortem completed | Debrief written in `docs/beta-reflection.md`

---

## 🔁 Feedback Loop

All beta findings will be stored in:

- `docs/beta-retro.md`  
- GitHub Issues: label = `beta-feedback`  
- Notion form (optional)

---

> Next:
> [→ 07-launch-plan.md](./07-launch-plan.md)  
> [→ 08-eol-plan.md](./08-eol-plan.md)
