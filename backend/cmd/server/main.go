package main

import (
	"context"
	"log"

	"myapp-backend/internal/app"
)

func main() {
	ctx := context.Background()

	application := app.New(ctx)

	if err := application.Run(); err != nil {
		log.Fatalf("app failed: %v", err)
	}
}
