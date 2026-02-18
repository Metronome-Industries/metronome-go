package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

func main() {
	withLoggingMiddleware()
	withMultipleMiddlewares()
}

// withLoggingMiddleware demonstrates using option.WithMiddleware to add a
// request/response logging interceptor. Middleware receives the outgoing
// request and a "next" function to forward it down the chain.
func withLoggingMiddleware() {
	logger := log.New(os.Stdout, "[metronome] ", log.LstdFlags)

	client := metronome.NewClient(
		option.WithMiddleware(func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
			// Before the request
			start := time.Now()
			logger.Printf("--> %s %s", req.Method, req.URL.Path)

			// Forward the request to the next handler
			res, err := next(req)

			// After the request
			duration := time.Since(start)
			if err != nil {
				logger.Printf("<-- %s %s ERROR (%s): %v", req.Method, req.URL.Path, duration, err)
			} else {
				logger.Printf("<-- %s %s %d (%s)", req.Method, req.URL.Path, res.StatusCode, duration)
			}

			return res, err
		}),
	)

	customers, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{})
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Printf("Found %d customers\n", len(customers.Data))
}

// withMultipleMiddlewares demonstrates chaining multiple middlewares.
// When multiple middlewares are provided, they are applied left to right.
// Middlewares set on the client run before middlewares set on individual requests.
func withMultipleMiddlewares() {
	// First middleware: add a custom request ID header
	addRequestID := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		req.Header.Set("X-Request-ID", "example-request-123")
		return next(req)
	}

	// Second middleware: log the request
	logRequest := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		fmt.Printf("Request ID: %s\n", req.Header.Get("X-Request-ID"))
		return next(req)
	}

	client := metronome.NewClient(
		// Middlewares are applied left to right: addRequestID runs first, then logRequest
		option.WithMiddleware(addRequestID, logRequest),
	)

	customers, err := client.V1.Customers.List(context.TODO(), metronome.V1CustomerListParams{})
	if err != nil {
		fmt.Printf("Error (expected in example): %s\n", err)
		return
	}
	fmt.Printf("Found %d customers\n", len(customers.Data))
}
