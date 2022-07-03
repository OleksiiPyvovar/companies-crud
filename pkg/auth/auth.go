package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type Interface interface {
	Validate(r *http.Request) error
}

type auth struct {
	secret string
}

func New(secret string) Interface {
	return &auth{secret: secret}
}

func (a *auth) Validate(r *http.Request) error {
	tokenString := getTokenFromRequest(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}

		return []byte(a.secret), nil
	})

	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func getTokenFromRequest(req *http.Request) string {
	token := strings.TrimSpace(req.Header.Get("Authorization"))
	if token == "" {
		token = req.URL.Query().Get("token")
	}

	return token
}
