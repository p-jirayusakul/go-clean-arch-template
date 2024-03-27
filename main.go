package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	handlers "github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/db"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/worker"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
	"golang.org/x/sync/errgroup"
)

const PATH_CONFIG = ".env"

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

var (
	cfg    = config.InitConfigs(PATH_CONFIG)
	dbConn = db.InitDatabase(cfg)
)

// @title           Clean Architecture
// @version         1.0
// @description     This is template clean arch

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host      localhost:3000
func main() {

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// plug database
	store := db.NewStore(dbConn)

	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.REDIS_ADDRESS,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	waitGroup, ctx := errgroup.WithContext(ctx)

	// plug controller
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// init validator
	app.Validator = middleware.NewCustomValidator()
	app.Use(middleware.ErrorHandler)

	// init log
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(middleware.LogHandler(logger))

	// add all plug to handler
	handlers.NewServerHttpHandler(app, &cfg, taskDistributor, store)

	// run all service
	runTaskProcessor(ctx, waitGroup, redisOpt, store)
	runServer(ctx, app, cfg, waitGroup)

	err := waitGroup.Wait()
	if err != nil {
		app.Logger.Fatal("error from wait group")
	}

}

func runTaskProcessor(
	ctx context.Context,
	waitGroup *errgroup.Group,
	redisOpt asynq.RedisClientOpt,
	store db.Store,
) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store)

	slog.Info("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		slog.Error("message: 'failed to start task processor'")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		slog.Info("graceful shutdown task processor")

		taskProcessor.Shutdown()
		slog.Info("task processor is stopped")

		return nil
	})
}

func runServer(ctx context.Context, app *echo.Echo, cfg config.Config, waitGroup *errgroup.Group) {

	waitGroup.Go(func() error {
		err := app.Start(":" + cfg.API_PORT)
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		slog.Info("graceful shutdown server")
		defer cancel()

		err := app.Shutdown(ctx)
		slog.Info("server is stopped")

		if err != nil {
			app.Logger.Fatal(err)
			fmt.Println("do this 2")

			return err
		}
		return nil
	})

}
