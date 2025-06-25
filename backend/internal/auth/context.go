package auth

import (
    "context"
    "net/http"
    "strings"
)

type contextKey string

const userIDKey contextKey = "userID"

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if strings.HasPrefix(authHeader, "Bearer ") {
            tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
            userID, err := ParseJWT(tokenStr)
            if err == nil {
                ctx := context.WithValue(r.Context(), userIDKey, userID)
                r = r.WithContext(ctx)
            }
        }
        next.ServeHTTP(w, r)
    })
}

func UserIDFromContext(ctx context.Context) (uint, bool) {
    userID, ok := ctx.Value(userIDKey).(uint)
    return userID, ok
}