package middlewares

import (
	"net/http"
)

var rolePermissions = map[string][]string{
	"user":  {"read"},
	"admin": {"read", "write", "delete"},
}

// RBACMiddleware ensures that the user has the required permission to access a route.
func RBACMiddleware(requiredPermission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate getting the role from a request header
			// In real-world scenarios, this could come from a token or session
			role := r.Header.Get("Role") // Example: "user" or "admin"

			// Check if the role exists in our rolePermissions map
			permissions, exists := rolePermissions[role]
			if !exists {
				http.Error(w, "Forbidden: role not found", http.StatusForbidden)
				return
			}

			// Check if the role has the required permission
			for _, perm := range permissions {
				if perm == requiredPermission {
					next.ServeHTTP(w, r)
					return
				}
			}

			// If no matching permission is found, deny access
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
		})
	}
}
