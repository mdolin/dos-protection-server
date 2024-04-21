// Description: This file contains the code for
// handling incoming HTTP requests.

package request

import (
	limiter "dos-protection-server/limiter"
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, rateLimiter *limiter.RateLimiter) {
	// Extract client ID from query parameters
	clientID := r.URL.Query().Get("clientId")
	if clientID == "" {
		http.Error(w, "Missing client ID parameter\n", http.StatusBadRequest)
		return
	}

	// Check if the request is allowed
	if rateLimiter.AllowRequest(clientID) {
		response := fmt.Sprintf("Request from client ID:%s allowed\n", clientID)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	} else {
		http.Error(w, "Request limit exceeded. Try again later\n", http.StatusTooManyRequests)
		return
	}
}
