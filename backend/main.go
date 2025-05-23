
package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"
    "operary/config"
    "operary/router"
)

func main() {
    // Load environment variables from .env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found. Continuing with system environment...")
    }

    // Initialize MongoDB
    config.InitMongo()

    // Start router
    r := router.NewRouter()
    log.Println("🚀 Operary API is running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
