package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Farindra968/go_project1/internal/config"
	health "github.com/Farindra968/go_project1/internal/http/handlers/healths"
	"github.com/Farindra968/go_project1/internal/http/handlers/students"
	"github.com/Farindra968/go_project1/internal/storage/sqlite"
)

func main() {
	fmt.Println("Loading configuration...")

	configCandidates := []string{
		"config/local.yml",
		"config/local.yaml",
		"../config/local.yml",
		"../config/local.yaml",
	}

	configPath := ""
	for _, candidate := range configCandidates {
		if _, err := os.Stat(candidate); err == nil {
			configPath = candidate
			break
		}
	}
	if configPath == "" {
		panic("config file not found in expected locations")
	}

	// Load configuration from YAML file and environment variables
	cfg, err := config.MustLoad(configPath)
	if err != nil {
		panic(err)
	}
	slog.Info("Configuration loaded successfully")

	// Database Setup
	storage, err := sqlite.NewSqLite(*cfg)
	if err != nil {
		panic(err)
	}

	slog.Info("Database connection established successfully")

	// Setup Router and HTTP server using the loaded configuration
	// The router is created using http.NewServeMux(), which is a multiplexer that matches incoming requests to their respective handlers based on the request URL.
	
	router:= http.NewServeMux()

	const apiPrefix = "/api/v1"
	router.HandleFunc("GET " + apiPrefix + "/health", health.HealthHandler())
	router.HandleFunc("POST " + apiPrefix + "/students", students.StudentHandler(storage))

	slog.Info("Server Starting successfully in:", "addr", cfg.HTTPServer.Addr)

	// Create an HTTP server using the loaded configuration, specifying the address and handler (router) for incoming requests.
	// The http.Server struct is used to configure the server's address and request handling behavior.
	server:= http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}

	// Create a channel to listen for OS signals for graceful shutdown
	// The done channel is used to signal when the server should shut down gracefully.
	done := make( chan os.Signal, 1)

	// Listen for OS signals to gracefully shutdown the server
	// Notify the done channel when an interrupt or termination signal is received
	// This allows the server to clean up resources and exit gracefully when the application is stopped.
	// The signal.Notify function registers the done channel to receive notifications of the specified signals.
	// The os.Interrupt signal is sent when the user presses Ctrl+C, and syscall.SIGINT and syscall.SIGTERM are common termination signals used in Unix-like operating systems.
	// syscall.SIGINT is the interrupt signal, and syscall.SIGTERM is the termination signal. When either of these signals is received, the done channel will be notified, allowing the application to perform any necessary cleanup before exiting.
	// This is a common pattern in Go applications to handle graceful shutdowns and ensure that resources are properly released before the application exits.
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the HTTP server in a separate goroutine to allow the main function to continue executing and listen for OS signals.
	// The server.ListenAndServe() method starts the HTTP server and blocks until the server is stopped or an error occurs.
	go func() {
			err :=server.ListenAndServe()
			slog.Error("Error starting server:", "err", err)

			if err != nil {
				slog.Error("Error starting server:", "err", err)
				panic(err)
			}
	} ()

	// Wait for a signal to gracefully shutdown the server
	// The <-done line blocks the main goroutine until a signal is received on the done channel, indicating that the server should shut down gracefully.
	<- done
	slog.Info("Server Stopping gracefully...")

	// Create a context with a timeout to allow the server to finish processing ongoing requests before shutting down.
	// The context.WithTimeout function creates a new context that will be canceled after the specified duration (5 seconds in this case).
	// This allows the server to complete any ongoing requests before shutting down, ensuring a graceful shutdown process.
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)

	defer cancel()

	// Shutdown the server gracefully, allowing it to finish processing ongoing requests before exiting.
	// The server.Shutdown method is called with the context created earlier, which will cancel the shutdown process if it takes longer than the specified timeout.
	// This ensures that the server has a chance to clean up resources and complete any ongoing requests before shutting down.

	server.Shutdown(ctx)

	slog.Info("Server Stopped gracefully")
	
	
}