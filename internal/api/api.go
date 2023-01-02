package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type API struct {
	server *http.Server
}

func (a *API) Start(ctx context.Context, wg *sync.WaitGroup) {
	// Critical to call wg.Done
	// defer ensures it is called, regardless of when the function returns
	defer wg.Done()

	// mux allows us to respond to different paths "/status" "/login" etc.
	mux := http.NewServeMux()
	mux.HandleFunc("/status", a.StatusHandler)

	// Need to create a server object so we can gracefully shutdown
	// http.ListenAndServe does not provide a way to call Shutdown
	a.server = &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	// Start serving in a new goroutine
	go a.serve(ctx)

	// Wait for the signal to shutdown
	<-ctx.Done()
	a.shutdown(ctx)
}

func (a *API) serve(ctx context.Context) {
	// Currently this does not properly respond to a failed server
	// The overall application would continue to run
	fmt.Println("API Server Listening")

	err := a.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		a.shutdown(ctx)
	}
}

func (a *API) shutdown(ctx context.Context) {
	fmt.Println("Shutting down API Server")

	// Creates a child context that has a 5 second timeout
	// This gives the server a hard deadline in the event of long requests
	shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	a.server.Shutdown(shutdownCtx)
}
