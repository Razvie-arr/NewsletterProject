package middleware

import (
	"errors"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	authSchema = "Bearer "
)

type Authenticator interface {
	VerifyToken(token string) (map[string]interface{}, error)
}

func NewAuthenticate(authenticator Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			header := request.Header.Get(authHeader)
			if header == "" {
				http.Error(responseWriter, "missing authentication header", http.StatusUnauthorized)
				return
			}

			token, err := parseBearerToken(header)
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
				return
			}

			claims, err := authenticator.VerifyToken(token)
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
				return
			}
			if err := verifyUserRole(claims); err != nil {
				http.Error(responseWriter, err.Error(), http.StatusForbidden)
				return
			}

			next.ServeHTTP(responseWriter, request)
		})
	}
}

func parseBearerToken(header string) (string, error) {
	if !strings.HasPrefix(header, authSchema) {
		return "", errors.New("invalid authentication schema")
	}
	token := strings.TrimPrefix(header, authSchema)
	if token == "" {
		return "", errors.New("the token is empty")
	}
	return token, nil
}

func verifyUserRole(claims map[string]interface{}) error {
	role, ok := claims["role"].(string)
	if !ok {
		return errors.New("missing user role claim")
	}
	if role == "authenticated" {
		return nil
	}
	return errors.New("invalid user role")
}
