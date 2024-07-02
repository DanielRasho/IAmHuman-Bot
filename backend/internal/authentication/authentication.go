package authentication

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_KEY"))

func CreateJWTToken(payload jwt.MapClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	s, err := t.SignedString(key)
	return s, err
}

func AuthenticateJWTToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString,
		// Defining function that returns wich key to use for the validation.
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil {
		return nil, err
	}

	// Getting the claims, if not raise an error.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token invalid")
	}

}
