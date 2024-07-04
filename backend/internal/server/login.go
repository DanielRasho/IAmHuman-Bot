package server

import (
	"backend/internal/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) handlePostLogin(c *gin.Context) {
	var credentials userCredentials

	// Retrieve body
	if err := c.ShouldBindBodyWithJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request params"})
		return
	}
	// Check if all params are filled.
	if credentials.Password == "" || credentials.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not empty string allowed"})
		return
	}

	var dbPassword string
	row := s.db.QueryRow("SELECT password from member WHERE username = $1", credentials.Username)

	// Check if user exists
	if err := row.Scan(&dbPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username or password does not match"})
		return
	}
	// Check if passwords are the same
	if dbPassword != credentials.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username or password does not match"})
		return
	}

	token, err := authentication.CreateJWTToken(jwt.MapClaims{"username": credentials.Username})

	// Check if jwt token creation was ok.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create JWT token"})
		return
	}

	// Write response
	response := gin.H{
		"message": "Login Sucessful.",
		"token":   token,
	}

	c.IndentedJSON(http.StatusOK, response)
}
