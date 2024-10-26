package main

import (
	"context"
	// "fmt"
	"log"
	// "memrizr/account/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	// initialize data sources
	ds, err := initDS()
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	router, err := inject(ds)
	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	// router := gin.Default()
	
	// handler.NewHandler(&handler.Config{
	// 	R: router,
	// })

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	
	// graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v \n", err)
		}
	}()

	log.Printf("Listening on port %v \n", srv.Addr)

	// wait for kill signal of channel
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// this blocks until a signal is passed into the quit channel
	<-quit

	// the context is used to inform the server it has 3 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown data sources
	if err := ds.close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}

	// shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v \n", err)
	}
}
