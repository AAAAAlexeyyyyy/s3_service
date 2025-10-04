package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AAAAAlexeyyyyy/s3_service/internal/infrastructure/config"
	"github.com/AAAAAlexeyyyyy/s3_service/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func Start() (*zap.SugaredLogger, error) {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	globConfig, err := config.LoadConfig()
	l := logger.LoadLogger(globConfig, os.Stdout)
	if err != nil {
		return l, fmt.Errorf("error loading config: %w", err)
	}

	return l, nil
}
