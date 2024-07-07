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

func (s *Server) Run() error {
	return s.engine.Run()
}

func NewServer() (*Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, err := database.New()
	
	router := gin.Default()
	router.Use(CORSMiddleware())

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

	assignRole := s.engine.Group("/assign-role")
	{
		assignRole.GET("/", s.handleGetAssignRole)
		assignRole.POST("/success", s.handlePostAssignRoleSuccess)
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
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }

        c.Next()
    }
}