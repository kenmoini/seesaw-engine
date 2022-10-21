package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kenmoini/seesaw-engine/internal/logging"
	"github.com/kenmoini/seesaw-engine/internal/server"
)

// RunHTTPServer will run the HTTP Server
func (config Config) RunHTTPServer() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.Seesaw.Server.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         config.Seesaw.Server.Host + ":" + config.Seesaw.Server.Port,
		Handler:      server.NewRouter(config.Seesaw.Server.BasePath),
		ReadTimeout:  config.Seesaw.Server.Timeout.Read * time.Second,
		WriteTimeout: config.Seesaw.Server.Timeout.Write * time.Second,
		IdleTimeout:  config.Seesaw.Server.Timeout.Idle * time.Second,
	}

	// Only listen on IPV4
	l, err := net.Listen("tcp4", config.Seesaw.Server.Host+":"+config.Seesaw.Server.Port)
	logging.Check(err, "Error listening on "+config.Seesaw.Server.Host+":"+config.Seesaw.Server.Port)

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the server is starting
	logging.LogStdOutInfo(fmt.Sprintf("Server is starting on %s\n", server.Addr))

	// Run the server on a new goroutine
	go func() {
		//if err := server.ListenAndServe(); err != nil {
		if err := server.Serve(l); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	logging.LogStdOutInfo(fmt.Sprintf("Server is shutting down due to %+v\n", interrupt))
	if err := server.Shutdown(ctx); err != nil {
		logging.LogStdErr(fmt.Sprintf("Server was unable to gracefully shutdown due to err: %+v", err))
	}
}
