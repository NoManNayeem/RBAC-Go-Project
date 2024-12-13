package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rbac-go/config"
	"rbac-go/internal/middlewares"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// LoginRequest represents the structure of the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the structure of the login response
type LoginResponse struct {
	Role string `json:"role"`
}

// loginHandler handles user login and assigns roles
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Assign roles based on username and password
	var role string
	if req.Username == "Nayeem" && req.Password == "Password" {
		role = "admin"
	} else {
		role = "public"
	}

	resp := LoginResponse{Role: role}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database
	if err := config.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer config.DB.Close()

	// Set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Public route, no permissions needed!")
	})
	mux.Handle("/admin", middlewares.RBACMiddleware("write")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Admin route, permission granted!")
	})))

	// CORS Middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // Local frontend
			"http://127.0.0.1:3000", // Local2 frontend
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Role"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS
	handler := corsMiddleware.Handler(mux)

	// Get the server port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Server running on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server stopped gracefully")
}
