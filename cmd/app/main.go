package main

import (
	"context"
	"projectsphere/cats-social/pkg/middleware/graceful"
	"projectsphere/cats-social/pkg/middleware/logger"
	"projectsphere/cats-social/pkg/protocol/httpListener"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger()

	httpProtocol := httpListener.Start()
	graceful.GracefulShutdown(
		context.TODO(),
		time.Duration(5*time.Second),
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return httpProtocol.Shutdown(ctx)
			},
		},
	)

	httpProtocol.Listen()
}
