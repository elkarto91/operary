# 🧮 Operary – Feature Prioritization

> “The biggest risk isn’t doing the wrong thing — it’s doing everything at once.”

---

## 🎯 Prioritization Frameworks Used

| Framework | Why Used |
|-----------|----------|
| **MoSCoW** | Helps define MVP cutoffs with clarity for v0.1  
| **RICE** | Used to validate long-term bets post-MVP  
| **Impact Mapping** (future) | To connect tasks to strategic goals

---

## ✅ MoSCoW Analysis (v0.1 – v1.0)

### 🟢 **Must Have**
| Feature | Notes |
|--------|-------|
| ✅ Create / Assign Task | Manual or API-based  
| ✅ Shift Start / Close | Core to shift continuity  
| ✅ Audit Log Engine | Immutable task + shift logs  
| ✅ Role-based API Tokens | Needed for multi-org demo  
| ✅ API Spec + Docs | Swagger + examples

---

### 🟡 **Should Have**
| Feature | Notes |
|--------|-------|
| 🔄 Webhook Input (from SCADA sim) | Simulated signal → task mapping  
| 🧾 Shift Summary Generator | Markdown → PDF archive  
| 📜 Commenting + Notes on Task | Enables rich audit  
| 📊 Minimal Metrics (task count, time open) | Internal dashboard/report

---

### 🟠 **Could Have**
| Feature | Notes |
|--------|-------|
| 📱 Terminal-style Frontend or CLI UI | For plant usage simulation  
| 🗂️ Archive Viewer | Access past shift reports  
| 🔄 Task Templates | Optional structured task types  
| 🧠 Tags + Metadata | For future smart querying

---

### 🔴 **Won’t Have (Yet)**
| Feature | Notes |
|--------|-------|
| 🧬 Full UI Frontend | Outside scope for backend-focused POC  
| 🌍 Multi-language | Optional after first pilot  
| 🤖 AI-generated tasks or summaries | Stretch goal (v1.1+)  
| 👥 User Signup/Login Flow | Admin token-controlled only for now

---

## 📊 RICE Score Highlights (Post-MVP Features)

| Feature | Reach | Impact | Confidence | Effort | RICE Score |
|--------|-------|--------|------------|--------|------------|
| External Webhook Input | 8 | 9 | 9 | 3 | **192**  
| Shift Summary Generator | 6 | 8 | 8 | 2 | **192**  
| Metrics Dashboard (Task Flow) | 6 | 7 | 8 | 2 | **168**  
| Commenting + Notes | 7 | 6 | 9 | 2 | **189**  
| Role Audit Viewer | 4 | 6 | 8 | 3 | 96  
| Task Tagging Engine | 3 | 6 | 7 | 2 | 73.5  
| Report Archive Browser | 5 | 5 | 6 | 3 | 100  

---

## 🧠 Prioritization Summary

1. Focus v0.1 on **core task lifecycle**, **audit trail**, and **shift boundaries**  
2. Use v0.2 to test **external ingestion** and **commentary-style rich logs**  
3. Keep v0.3 focused on **user feedback**, tagging, PDF exports  
4. Push dashboarding + metrics post-pilot or v1.1

---

> Next:
> [→ 06-beta-plan.md](./06-beta-plan.md)  
> [→ 07-launch-plan.md](./07-launch-plan.md)  
> [→ 08-eol-plan.md](./08-eol-plan.md)
