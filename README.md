## DoS Protection Server

This project implements a simple HTTP Denial-of-Service (DoS) protection system consisting of a client and server in Go language. The client sends HTTP requests to the server, and the server enforces rate limiting based on client IDs to prevent DoS attacks.

### Main bits of the project
* Main File
* Ratelimiter Package
* Request Package

### Structure of the project
```
.
├── README.md
├── go.mod
├── go.sum
├── limiter
│   ├── ratelimiter.go
│   └── ratelimiter_test.go
├── main.go
└── request
    └── request.go
```

### Features
* Implements an HTTP server using Gorilla mux.
* Enforces rate limiting per client based on incoming requests.

### Server Implementation Details

* The server listens on port 8080 and exposes an endpoint (/) that checks rate limits based on client IDs.
* Rate limiting logic:
  - Each client is limited to a maximum of 5 requests within 5 seconds.
  - The time window starts from the client's first request and resets after 5 seconds.


### Prerequisites

- Go (Golang) installed on your system.
- Dependencies managed using Go modules.

### Setting Up and Running the Server

1. Clone the repository:

   ```shell
   git clone <repository-url>

2. Navigate to the server directory:

   ```shell
   go mod download

3. Build and run the server:
   ```shell
   go build
   ./dos-protection-server

4. The server will start listening on http://localhost:8080.

### Sending Requests Using Curl
1. To test the server with curl:
   ```shell
   curl http://localhost:8080/?clientId=1

### Configuration
* Adjust the rate limiting parameters (limit and window duration) in ratelimiter.go as per your requirements.