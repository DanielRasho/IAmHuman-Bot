package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type discordCode string

func (s *Server) handleGetAssignRole(c *gin.Context) {

	invitationId := c.Query("invitation_id")

	if invitationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invitation_id query must not be empty"})
		return
	}

	row := s.db.QueryRow(
		`select 
			id,
			class,
			section ,
			missing_uses ,
			created_at ,
			server_id ,
			role_id 
		from invitation
		where id = $1 ;`,
		invitationId)

	var invitationData invitation
	// Check if user exists
	if err := row.Scan(
		&invitationData.Id,
		&invitationData.Class,
		&invitationData.Section,
		&invitationData.MissingUses,
		&invitationData.CreateAt,
		&invitationData.ServerId,
		&invitationData.RoleId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invitation does not exist on DB"})
		return
	}

	oauthURL := os.Getenv("DB_DATABASE")

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":         "invitation data fetched succesfully",
		"invitation_data": invitationData,
		"Oauth":           oauthURL,
	})
}

func (s *Server) handlePostAssignRoleSuccess(c *gin.Context) {

}
