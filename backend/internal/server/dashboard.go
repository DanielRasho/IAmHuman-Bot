package server

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
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

	// FETCH INVITATIONS
	base_url := os.Getenv("CLIENT_URL") + "assign-role/"

	rows, err := s.db.Query(
		`select 
			id,
			class,
			section ,
			missing_uses ,
			to_char(created_at, 'DD-MM-YYYY'),
			server_id ,
			role_id,
			concat($1::text,id) as invitation_url
		from invitation
		order by class, section
		limit $2 offset $3;`,
		base_url, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data from DB " + err.Error()})
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
			&newInvitation.RoleId,
			&newInvitation.InvitationURL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot scan rows from DB"})
			return
		} else {
			invitations = append(invitations, newInvitation)
		}
	}

	// COUNTING TOTAL PAGES
	var rowsCountNum int
	rowsCount := s.db.QueryRow("select count(id) from invitation;")
	if err := rowsCount.Scan(&rowsCountNum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data from DB " + err.Error()})
		return
	}

	total_pages := int(math.Ceil(float64(rowsCountNum) / float64(limit)))

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":      "Data succesfully fetched.",
		"total_pages":  total_pages,
		"current_page": page,
		"items":        invitations,
	})
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
