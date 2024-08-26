package server

import (
	"fmt"
	"github.com/kirkegaard/go-reverse-proxy/internal/configs"
	"net/http"
	"net/url"
)

// Run start the server on defined port
func Run() error {
	// Load configurations from config file
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("could not load configuration: %v", err)
	}

	// Creates a new router
	mux := http.NewServeMux()

	// Register health check endpoint
	mux.HandleFunc("/ping", ping)

	// Iterating through the configuration resource and registering them
	// into the router.
	for _, resource := range config.Resources {
		fmt.Printf("Registering resource: %s to %s \n", resource.Endpoint, resource.Destination)

		url, _ := url.Parse(resource.Destination)
		proxy := NewProxy(url)

		handler := ProxyRequestHandler(proxy, url, resource.Endpoint)
		mux.Handle(resource.Endpoint, http.StripPrefix(resource.Endpoint, http.HandlerFunc(handler)))
	}

	// Running proxy server
	fmt.Printf("Server is running on %s:%s\n", config.Host, config.Port)
	if err := http.ListenAndServe(config.Host+":"+config.Port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}

	return nil
}
