
# 🛎️ Notification Services

This folder contains background and auxiliary services that operate outside the main request/response cycle.

---

## `notification.go`

### Purpose:
A long-running background dispatcher that simulates future integration points with external systems, such as:

- Webhook pings
- Event streaming (Kafka, NATS, Redis Streams)
- Retry queue for failed deliveries
- Scheduled job execution

### How It Works:
- Runs as a goroutine in `main.go`
- Logs heartbeat every 15 seconds using Zap
- Prints idle queue status with timestamps

### Future Ideas:
- Use Mongo change streams to detect new audit/task events
- Push status updates to external ERP/MES or reporting systems
- Integrate with `OpenTelemetry` for trace-spanning events

---

## 📦 Usage in Main

```go
import "operary/internal/services"

...

services.StartNotificationService(logger)
```

This ensures the dispatcher boots with your app and provides observability immediately.

## Development Notes

These helpers are intentionally lightweight. Expand the dispatcher with message queues or extra logging as your experiments evolve.
