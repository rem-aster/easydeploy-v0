package main

import (
	"context"
	"log"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
