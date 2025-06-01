package main

import (
	"net/http"

	"github.com/elkarto91/operary/config"
	"github.com/elkarto91/operary/internal/corepad"
	"github.com/elkarto91/operary/internal/opsmirror"
	"github.com/elkarto91/operary/internal/services"
	"github.com/elkarto91/operary/router"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
	db := config.GetMongoDB()

	opsmirror.Init(db)
	corepad.Init(db)

	services.StartNotificationService(sugar)

	sugar.Info("📡 Starting Operary API on :8080")
	r := router.NewRouterWithLogger(sugar)

	http.ListenAndServe(":8080", r)
}
