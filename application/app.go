package application

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/redis/go-redis/v9"
)

type App struct {
	rdb    *redis.Client
	server *fiber.App
}

func New() *App {

	redis_addr := os.Getenv("REDIS_ADDR")

	if redis_addr == "" {
		redis_addr = "localhost:6379"
	}
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: redis_addr,
		}),
		server: fiber.New(),
	}
	app.server.Use(logger.New())
	app.LoadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis")
		}
	}()

	fmt.Println("starting server")

	ch := make(chan error, 1)

	go func() {
		err = a.server.Listen(":4000")
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)

	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		fmt.Println("server shuting down")
		return a.server.ShutdownWithTimeout(time.Second * 10)
	}

}
