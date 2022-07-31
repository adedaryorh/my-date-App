package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenService interface {
	GenerateToken(Claims) (string, error)
	ValidateToken(string) error
}

type Claims struct {
	Name  string
	Email string
}

type jWTTokenService struct {
	hMacSecret []byte
}

func NewTokenService(secret string) jWTTokenService {
	return jWTTokenService{
		hMacSecret: []byte(secret),
	}
}

func (ts jWTTokenService) GenerateToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  c.Name,
		"email": c.Email,
		"nbf":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString(ts.hMacSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts jWTTokenService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return ts.hMacSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return err
}
