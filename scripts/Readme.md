# 🧪 Operary – Shift Simulation Scripts

This directory contains scripts to test core Operary flows using the public REST API.

---

## 📋 What’s Included

| Script | Description |
|--------|-------------|
| `sim_shift_test.go` | Starts a shift, creates tasks, updates statuses, and closes the shift — all via API calls
| `sample_media_upload.go` | Uploads a note with example media attachment
| *(future)* | Add escalation test, webhook trigger test, or error response test

---

## 🚀 Prerequisites

- Go 1.18+ installed  
- Operary backend running locally on `http://localhost:8080/v1`  
- Valid `X-Org-Token` set in the script

---

## 🧠 How It Works

This simulation performs:

1. Starts a new shift as a supervisor  
2. Creates 2–3 operational tasks  
3. Simulates closing, escalating, and updating tasks  
4. Closes the shift  
5. Confirms that actions are logged and printed

---

## 🧪 Run the Script

```bash
cd scripts/
go run sim_shift_test.go
go run sample_media_upload.go # send a note with media


Expected Output

🕓 Shift started: 79fa...
📋 Task created: Check pump pressure
🔄 Task ... → in_progress
🔄 Task ... → closed
🔄 Task ... → escalated
🔒 Shift ... closed.
✅ Simulated shift complete.
