package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/zengineeringgroup/golang-service/internal/busy"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	done := make(chan error, 1)

	go run(ctx, done)

	err := <-done
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, done chan error) {
	done <- busy.Busy(ctx)
}
