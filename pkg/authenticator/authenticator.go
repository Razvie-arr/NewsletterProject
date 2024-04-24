package authenticator

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret string
}

func NewJWTAuthenticator(secret string) JWTAuthenticator {
	return JWTAuthenticator{secret: secret}
}

func (authenticator JWTAuthenticator) VerifyToken(token string) (map[string]interface{}, error) {
	var claims jwtClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(authenticator.secret), nil
	})
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"role": claims.Role,
		"uuid": claims.Uuid,
	}
	return result, nil
}

type jwtClaims struct {
	jwt.RegisteredClaims

	Uuid string `json:"sub"`
	Role string `json:"role"`
}
