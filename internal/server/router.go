package server

import (
	"fmt"
	"net/http"

	"github.com/kenmoini/seesaw-server/internal/config"
	"github.com/kenmoini/seesaw-server/internal/logging"
)

// NewRouter generates the router used in the HTTP Server
func NewRouter(basePath string) *http.ServeMux {
	if basePath == "" {
		basePath = "/" + config.DefaultBasePath
	}
	// Create router and define routes and return that router
	router := http.NewServeMux()

	// Version Output
	router.HandleFunc(basePath+"/version", func(w http.ResponseWriter, r *http.Request) {
		logging.LogNetworkRequestStdOut(r.Method+" "+basePath+"/version", r)
		fmt.Fprintf(w, config.AppName+" version: %s\n", config.AppVersion)
	})

	// Healthz endpoint for kubernetes platforms
	router.HandleFunc(basePath+"/healthz", func(w http.ResponseWriter, r *http.Request) {
		logging.LogNetworkRequestStdOut(r.Method+" "+basePath+"/healthz", r)
		fmt.Fprintf(w, "OK")
	})

	return router
}
