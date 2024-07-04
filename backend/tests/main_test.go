package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load environment variables from .env file
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("No .env file found")
	}

	// Run the tests
	code := m.Run()

	// Exit with the exit code from m.Run()
	os.Exit(code)
}
