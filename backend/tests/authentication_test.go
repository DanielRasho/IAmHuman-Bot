package tests

import (
	"backend/internal/authentication"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestAuthentication(t *testing.T) {
	var token string
	t.Run("creating Token", func(t *testing.T) {
		got, err := authentication.CreateJWTToken(jwt.MapClaims{
			"username": "Daniel",
			"book":     "harry poter",
		})
		if err != nil {
			t.Errorf("Failed to create token: %v", err)
		}
		token = got
		fmt.Println(token)
	})

	t.Run("Verifying Token", func(t *testing.T) {
		got, err := authentication.AuthenticateJWTToken(token)
		if err != nil {
			t.Errorf("Failed to authenticate token: %v", err)
		}

		// Iterate through the claims
		for key, value := range got {
			fmt.Printf("%s: %v\n", key, value)
		}
	})
	t.Run("Verifying Invalid Token", func(t *testing.T) {
		_, err := authentication.AuthenticateJWTToken(token + "a")
		if err == nil {
			t.Errorf("Token should be invalid: %v", err)
		}
	})
}
