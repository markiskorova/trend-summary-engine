package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Replace with os.Getenv(...) in production

// GenerateJWT creates a signed token with the user's ID.
func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtSecret)
}

// ParseJWT validates and parses a token, returning the user ID.
func ParseJWT(tokenStr string) (uint, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid claims")
    }

    idFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("user_id not found in token")
    }

    return uint(idFloat), nil
}