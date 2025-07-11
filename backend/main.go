package main

import (
	"net/http"

	"github.com/elkarto91/operary/config"
	"github.com/elkarto91/operary/internal/auditsync"
	"github.com/elkarto91/operary/internal/authz"
	"github.com/elkarto91/operary/internal/configmgr"
	"github.com/elkarto91/operary/internal/corepad"
	"github.com/elkarto91/operary/internal/equiptrust"
	"github.com/elkarto91/operary/internal/eventbus"
	"github.com/elkarto91/operary/internal/flowgrid"
	"github.com/elkarto91/operary/internal/integrations"
	"github.com/elkarto91/operary/internal/notifier"
	"github.com/elkarto91/operary/internal/opsmirror"
	"github.com/elkarto91/operary/internal/permitgrid"
	"github.com/elkarto91/operary/internal/search"
	"github.com/elkarto91/operary/internal/sensorvault"
	"github.com/elkarto91/operary/internal/services"
	"github.com/elkarto91/operary/internal/supplymesh"
	"github.com/elkarto91/operary/internal/traceboard"
	"github.com/elkarto91/operary/internal/trainops"
	"github.com/elkarto91/operary/internal/twinboard"
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
	auditsync.Init(db)
	authz.Init(db)
	permitgrid.Init(db) // ✅ NEW
	flowgrid.Init(db)
	traceboard.Init(db)
	equiptrust.Init(db)
	sensorvault.Init(db)
	trainops.Init(db)
	twinboard.Init(db)
	supplymesh.Init(db)
	notifier.Init(db)
	eventbus.Init(db)
	configmgr.Init(db)
	search.Init(db)
	integrations.Init(db)
	services.StartNotificationService(sugar)

	sugar.Info("📡 Starting Operary API on :8080")
	r := router.NewRouterWithLogger(sugar)

	if err := http.ListenAndServe(":8080", r); err != nil {
		sugar.Fatalw("failed to start server", "error", err)
	}
}
