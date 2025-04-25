package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lxnewayfarer/shortlinks/routes"
	"log/slog"
	"net/http"
	"os"
	"context"
	"os/signal"
	"time"
)

func setupEnvironment() error {
	requiredEnvVars := []string{"APP_URL", "PORT", "REDIS_URL"}

	if err := godotenv.Load(); err != nil {
		return err
	}

	for _, x := range requiredEnvVars {
		_, exist := os.LookupEnv(x)
		if !exist {
			return fmt.Errorf("required variable %s is missing", x)
		}
	}

	return nil
}

func main() {
	slog.Info("Start Shortlinks")

	if err := setupEnvironment(); err != nil {
		slog.Error("Failed to setup environment", "err", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	slog.Info("Starting server",
		"host", "localhost",
		"port", port,
	)
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           routes.Init(),
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		slog.Info("Server is running", "address", "http://localhost:"+port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "err", err)
			os.Exit(1)
		}
	}()
	<-ctx.Done()
	slog.Info("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("Server shutdown error", "err", err)
	}
	slog.Info("Goodbye!")
}
