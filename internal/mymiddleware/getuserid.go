package mymiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xtasysensei/go-poll/internal/config"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

var jwtKey = []byte(config.Envs.JWTSecret)

func WithUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Split the "Bearer" part from the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		// Parse the JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Extract the user ID from the custom claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var userID int
			switch v := claims["userID"].(type) {
			case string:
				userID, err = strconv.Atoi(v)
				if err != nil {
					http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
					return
				}
			case float64: // JWT numbers are float64 by default
				userID = int(v)
			default:
				http.Error(w, "userID claim is missing or not a valid type", http.StatusUnauthorized)
				return
			}

			// Add the user ID to the request context

			ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
	})
}
