
package services

import (
    "time"

    "go.uber.org/zap"
)

func StartNotificationService(logger *zap.SugaredLogger) {
    go func() {
        ticker := time.NewTicker(15 * time.Second)
        defer ticker.Stop()

        for {
            <-ticker.C
            logger.Infow("🔔 Notification dispatcher active",
                "timestamp", time.Now().Format(time.RFC3339),
                "status", "idle",
                "queue", "0 events",
            )
        }
    }()
}
