package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/lxnewayfarer/shortlinks/routes"
	"github.com/lxnewayfarer/shortlinks/storage"
)

func main() {
	slog.Info("Start Shortlinks")

	if err := setupEnvironment(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	rdb, err := storage.InitRedis()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := createServer(port, rdb)

	go func() {
		slog.Info("Starting server", "address", "http://localhost:"+port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-ctx.Done()

	shutdown(server)
}

func shutdown(server *http.Server) {
	slog.Info("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	slog.Info("Goodbye!")
}

func createServer(port string, rdb storage.RedisClient) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           routes.Init(rdb),
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
}

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
