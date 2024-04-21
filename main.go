// Description: A simple web server that listens on port 8080 and
// responds to all requests with a message.

package main

import (
	limiter "dos-protection-server/limiter"
	request "dos-protection-server/request"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the rate limiter
	rateLimiter := limiter.NewRateLimiter(5, 5*time.Second)

	// Create a new router instance
	router := mux.NewRouter()

	// Define the handler function for incoming requests
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request.HandleRequest(w, r, rateLimiter)
	})

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
