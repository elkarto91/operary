# 🗺️ Operary – Product Roadmap

> “The roadmap is not a promise. It’s a narrative of bets we’re willing to place — with clarity.”

---

## 🛠️ v0.1 – Prototype (Proof of Coordination)

📅 Timeline: Month 1–2  
🎯 Goal: Prove end-to-end signal → task → log workflow  
🏁 Outcome: Code repo, test API, documented trace flow

### Milestones
- ✅ Architecture design finalized (docs/architecture.md)
- ✅ OpenAPI spec defined (api-spec/openapi.yaml)
- ✅ Mongo + GoLang backend initialized
- 🔄 Task CRUD, Shift CRUD APIs
- 🔄 Minimal CLI/Postman testing suite
- 🔄 Audit log recording logic (mutable → sealed)
- 🔄 Sample shift report generator (Markdown/PDF)
- 🔄 Excalidraw system & wireframes

---

## 🧪 v0.2 – Beta Pilotable Version

📅 Timeline: Month 3–4  
🎯 Goal: Deployable on internal VM or edge server  
🏁 Outcome: Real data captured via APIs

### Features
- 🧩 Basic token-based Org setup + access control  
- 📦 External system webhook receiver (simulated)  
- 📲 Mobile-first JSON UI or TUI (optional)  
- 🗃️ Local persistence with TTL & audit locks  
- 🧪 Structured test cases + coverage report  
- 📑 Beta onboarding README + usage examples

---

## 🌍 v0.3 – Field Pilot Ready (External Validation)

📅 Timeline: Month 5–6  
🎯 Goal: Trial at a real industrial site or with a simulated partner org  
🏁 Outcome: Measurable coordination loop feedback

### Deliverables
- 📡 Real SCADA-like input simulation  
- 📘 Role-based view UX demo  
- 🔐 Token management + limited ACL  
- 📦 Archiving of past shift reports  
- 🔁 Shift schedule simulator (scripted)  
- 🧠 Interview questions for pilot feedback  
- 📈 Internal report: “Did Operary reduce decision time?”

---

## 🚀 v1.0 – Go-to-Market (GTM Ready)

📅 Timeline: Month 7–9  
🎯 Goal: Public page + pitch + first proof client

### Components
- 🌐 Launch on [systemsignal.dev/projects](https://systemsignal.dev/projects)  
- 🧩 Integrate with real data tool (PostHog, InfluxDB, SCADA simulation)  
- 📄 Create PDF + Notion deck for investor/advisor conversations  
- 🎥 Record walk-through of system  
- 🎤 Author thought-leadership post: "Coordination-as-Infrastructure"  
- 🔁 Inbound flow: calendly bookings + demo requests

---

## 🧭 Strategic Track (Parallel)

| Phase | Strategic Outcome |
|-------|-------------------|
| v0.1 | Product thesis validation  
| v0.2 | Internal pilot credibility  
| v0.3 | Field partner trust loop  
| v1.0 | Market signal + pitch readiness  
| v1.1+ | Consulting client install, or FOSS wrapper (opt-in)

---

## 💬 Meta Notes

- Roadmap is **narrative, not contract**  
- Tasks are tracked in `/tasks` or GitHub Projects  
- Feedback loops from early testers will shape new bets  
- This doc evolves continuously as new clarity emerges

---

> Next up:
> [→ 05-prioritization.md](./05-prioritization.md)  
> [→ 06-beta-plan.md](./06-beta-plan.md)  
> [→ 07-launch-plan.md](./07-launch-plan.md)
