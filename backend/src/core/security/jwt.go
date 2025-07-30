package security

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func getSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodRS256
}

type ValidationConfigRepo interface {
	GetPublicKey() *rsa.PublicKey
}

func CreateJWT(privateKey crypto.Signer, id string) (string, error) {
	now := time.Now()
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 30)),
		},
	}
	token := jwt.NewWithClaims(getSigningMethod(), claims)
	return token.SignedString(privateKey)
}

func DecodeAndValidateJWT(publicKey crypto.PublicKey, tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != getSigningMethod().Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an RSA public key type: %T", publicKey)
		}
		return rsaPublicKey, nil
	})
	if err != nil {
		slog.Error("Token parsing failed", "reason", err)
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}
	if !token.Valid {
		slog.Error("Invalid token")
		return nil, fmt.Errorf("token is invalid")
	}
	parsedClaims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}
	return parsedClaims, nil
}
