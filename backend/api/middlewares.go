package api

import (
	"context"
	"net/http"
	"todo_app_backend/internal/app/utils"
)

type UserKey string

const (
	UserIDKey UserKey = "userID"
)

func RespondJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func UserOnlyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check for token in request header
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) <= len("Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr = tokenStr[len("Bearer "):]
		payload, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, payload.UserID)
		r = r.WithContext(ctx)

		// call next handler
		h.ServeHTTP(w, r)
	})
}
