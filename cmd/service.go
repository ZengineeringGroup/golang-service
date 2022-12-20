package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/zengineeringgroup/golang-service/internal/api"
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go busy.Busy(ctx, wg)

	a := api.API{}
	wg.Add(1)
	go a.Start(ctx, wg)

	wg.Wait()
	done <- nil
}
