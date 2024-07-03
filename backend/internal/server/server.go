package server

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"backend/internal/authentication"
	"backend/internal/database"
)

type Server struct {
	port int

	db database.Service

	engine *gin.Engine
}

func (s *Server) Run() error {
	return s.engine.Run()
}

func NewServer() (*Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	router := gin.Default()
	db, err := database.New()

	// Check if connection to DB was succesfull
	if err != nil {
		return nil, err
	}

	NewServer := &Server{
		port:   port,
		db:     db,
		engine: router,
	}

	NewServer.registerRoutes()

	return NewServer, nil
}

func (s *Server) registerRoutes() {
	login := s.engine.Group("/login")
	{
		login.POST("/", s.handlePostLogin)
	}

	dashboard := s.engine.Group("/dashboard")
	dashboard.Use(AuthRequiredMiddleware())
	{
		dashboard.GET("/", s.handleGetDashboard)
		dashboard.POST("/", s.handlePostDashboard)
		dashboard.DELETE("/", s.handleDeleteDashboard)
	}
}

func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		_, err := authentication.AuthenticateJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
