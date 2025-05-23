# 🛑 Operary – End-of-Life (EOL) Plan

> “Even well-built systems need graceful exits. Sunsetting is not an afterthought — it’s part of product integrity.”

---

## 🎯 Why Have an EOL Plan?

In enterprise settings, unmaintained software becomes:

- A **liability** to security, compliance, and audit  
- A **risk** to teams depending on workflows  
- A **trust gap** when deprecation isn't communicated clearly

**Operary's EOL plan** ensures:

✅ Responsible withdrawal  
✅ Data handover or export  
✅ Minimal disruption  
✅ Continued trust

---

## 🕰️ Triggers for Initiating EOL

| Trigger | Description |
|---------|-------------|
| Business exit | Product no longer aligns with strategic goals  
| Replacement | Internal tool supersedes Operary  
| Acquisition | IP or client shifted to new ecosystem  
| Community stall | Usage or support falls below active threshold  
| Internal decision | Maintainer resources reallocated

---

## 📅 EOL Timeline (Suggested Policy)

| Phase | Duration | Activities |
|-------|----------|------------|
| ⚠️ Notice Phase | 60–90 days |  
- Email & publish EOL announcement  
- Freeze new signups or org creation  
- Begin export capability testing  
| 🛠️ Transition Phase | 30–60 days |  
- Support data export  
- Prepare backup archive  
- Disable write actions, keep logs accessible  
| 🔒 Sunset Phase | Day 91+ |  
- Shut down endpoints  
- Archive GitHub, mark deprecated  
- Redirect project URLs to EOL notice

---

## 📦 Data Export Plan

| Artifact | Export Format |
|----------|---------------|
| Tasks & shifts | JSON or CSV bundle  
| Audit logs | JSON (per entity or date-range)  
| Org metadata | YAML / JSON  
| API usage stats | Optional CSV or Grafana export

Export script will live in:  
`/scripts/export_eol_package.go`

---

## 📣 EOL Communication Plan

| Channel | Action |
|---------|--------|
| GitHub | Add `ARCHIVED.md` and repo banner  
| SystemSignal.dev | Add EOL blog entry + link from `/projects`  
| Email | Notify beta users + collaborators  
| README | Move project to **"historical"** portfolio section  
| LinkedIn | Post closure reflections, learning, reuse notes

---

## ✅ Maintainer Guarantees

- Read-only logs will remain accessible for 6 months after EOL  
- GitHub repo and API spec will remain public for educational reuse  
- Final PDF package and API freeze will be documented in `docs/eol/`

---

## 🧠 Reflections After EOL

If EOL is reached:

- Publish: `/docs/what-we-learned.md`  
- Archive: systemsignal.dev/operary-eol  
- Convert: elements into open-source templates (task engine, audit core)  
- Retire: API keys, tokens, production databases securely

---

> Sunset well, and users will return for your sunrise.

