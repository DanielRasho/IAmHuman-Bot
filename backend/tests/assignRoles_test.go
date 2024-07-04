package tests

import (
	"backend/internal/server"
	"fmt"
	"testing"
)

func TestAssignRoles(t *testing.T) {
	t.Run("Exchange Discord code for access token", func(t *testing.T) {
		got, err := server.ExchangeCodeForToken("ff4f98jvlHES1aKEfCkzJoYfEvNkmz")
		if err != nil {
			t.Errorf("Failed to create token: %v", err)
		}
		fmt.Printf("AccessToken : %s\n", got.AccessToken)
	})
}
