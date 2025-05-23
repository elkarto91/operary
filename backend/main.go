
package main

import (
    "net/http"

    "github.com/joho/godotenv"
    "go.uber.org/zap"
    "operary/config"
    "operary/router"
    "operary/internal/services"

)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    sugar := logger.Sugar()

    sugar.Info("🌐 Loading environment...")
    if err := godotenv.Load(); err != nil {
        sugar.Warn("No .env file found. Using system environment.")
    }

    sugar.Info("🔌 Connecting to MongoDB...")
    config.InitMongo()

    services.StartNotificationService(sugar)

    sugar.Info("📡 Starting Operary API on :8080")
    r := router.NewRouterWithLogger(sugar)

    http.ListenAndServe(":8080", r)
}
