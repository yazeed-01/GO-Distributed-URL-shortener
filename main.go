package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"urlShorter/initializers"
	"urlShorter/routes"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	// stop server 1
	stopServer1 := make(chan struct{})

	// prepare server 1
	srv1 := &http.Server{
		Addr:    ":8080",
		Handler: logRequests(routes.SetupRoutes(), "Server 1 (Port 8080)"),
	}

	go func() {
		log.Println("Server 1 is running on port 8080")
		if err := srv1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to run server 1: %v", err)
		}
	}()

	// prepare server 2
	srv2 := &http.Server{
		Addr:    ":8081",
		Handler: logRequests(routes.SetupRoutes(), "Server 2 (Port 8081)"),
	}

	go func() {
		log.Println("Server 2 is running on port 8081")
		if err := srv2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to run server 2: %v", err)
		}
	}()

	// timer to stop server 1
	go func() {
		time.Sleep(10 * time.Hour)
		close(stopServer1)
	}()

	<-stopServer1
	log.Println("Stopping Server 1...")

	// stop server 1
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv1.Shutdown(ctx); err != nil {
		log.Fatalf("Server 1 forced to shutdown: %v", err)
	}

	log.Println("Server 1 has been stopped.")
	select {}
}

func logRequests(h http.Handler, serverName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s received request from %s\n", serverName, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}
