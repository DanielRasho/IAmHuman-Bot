package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetDashboard(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	rows, err := s.db.Query(
		`select 
			id,
			class,
			section ,
			missing_uses ,
			created_at ,
			server_id ,
			role_id 
		from invitation limit $1 offset $2;`,
		limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data from DB"})
		return
	}
	defer rows.Close()

	invitations := []invitation{}
	for rows.Next() {
		var newInvitation invitation
		err := rows.Scan(
			&newInvitation.Id,
			&newInvitation.Class,
			&newInvitation.Section,
			&newInvitation.MissingUses,
			&newInvitation.CreateAt,
			&newInvitation.ServerId,
			&newInvitation.RoleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot scan rows from DB"})
			return
		} else {
			invitations = append(invitations, newInvitation)
		}
	}

	c.IndentedJSON(http.StatusOK, invitations)
}

func (s *Server) handlePostDashboard(c *gin.Context) {
	var invite invitation

	if err := json.NewDecoder(c.Request.Body).Decode(&invite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	_, err := s.db.Exec(`INSERT into invitation 
		(class, section, missing_uses, server_id, role_id)
		VALUES
		($1, $2, $3, $4, $5); `,
		invite.Class,
		invite.Section,
		invite.MissingUses,
		invite.ServerId,
		invite.RoleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cant create invitation."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "invitation created."})
}

func (s *Server) handleDeleteDashboard(c *gin.Context) {
	invitationId, err := strconv.Atoi(c.Query("invitation_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invitation_id query must not be empty."})
		return
	}

	_, err = s.db.Exec(`DELETE from invitation 
		WHERE id = $1`,
		invitationId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "If ID was valid, the invitation got deleted."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "invitation deleted."})
}
