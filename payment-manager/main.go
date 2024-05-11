package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"payment-manager/config/db"
	"payment-manager/config/inits"
	"payment-manager/helper"
	"payment-manager/router"
	"syscall"
	"time"
)

func main() {
	var isDev bool
	if helper.MyConfig.Environment == "PROD" {
		isDev = false
	} else {
		isDev = true
	}

	r := router.CreateRouter(isDev)
	r.GET("/", inits.HandlerHealthCheck)

	router.InitRoute(r)

	server := &http.Server{
		Addr:         helper.MyConfig.ServerPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit

		db.CloseConnectionDB()

		helper.Log.Println(helper.MyConfig.AppName, " is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			helper.Log.Fatalf("Could not gracefully shutdown the server %v: %v\n", helper.MyConfig.AppName, err)
		}
		close(done)
	}()

	helper.Log.Println(helper.MyConfig.AppName, "is ready to handle request at", helper.MyConfig.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		helper.Log.Fatalf("Could not listen on %s: %v\n", helper.MyConfig.ServerPort, err)
	}

	<-done
	helper.Log.Println(helper.MyConfig.AppName, " stopped")
}
