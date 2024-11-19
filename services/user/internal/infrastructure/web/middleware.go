package web

import (
	"context"
	"localloop/services/user/internal/domain/user"
	"net/http"
)

type Authenticator struct {
	userService *user.Service
}

func NewAuthenticator(userService *user.Service) *Authenticator {
	return &Authenticator{
		userService: userService,
	}
}

func (a *Authenticator) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized - No token provided", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// Validate token and get claims
		claims, err := a.userService.ValidateToken(token)
		if err != nil {
			http.Error(w, "Unauthorized - "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Verify user exists in database
		user, err := a.userService.Get(claims.Email)
		if err != nil {
			http.Error(w, "Unauthorized - User not found", http.StatusUnauthorized)
			return
		}

		// Add user context and proceed
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
