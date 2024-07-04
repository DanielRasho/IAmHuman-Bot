package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const DISCORD_API_GET_TOKEN string = "https://discord.com/api/oauth2/token"

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

	oauthURL := os.Getenv("DISCORD_OAUTH_URL") + "&state=" + invitationId

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":         "invitation data fetched succesfully",
		"invitation_data": invitationData,
		"Oauth":           oauthURL,
	})
}

func (s *Server) handlePostAssignRoleSuccess(c *gin.Context) {

	var discordCode discordCode
	if err := c.ShouldBindJSON(&discordCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Discord code is in bad format"})
		return
	} else if discordCode.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Discord code cannot be an empty string"})
		return
	}

	// Fetch url info
	var invitation invitation
	row := s.db.QueryRow(
		`select 
			id,
			class,
			section ,
			missing_uses ,
			created_at ,
			server_id ,
			role_id 
		from invitation where id = $1;`,
		discordCode.InvitationId)

	// Check if user exists
	if err := row.Scan(
		&invitation.Id,
		&invitation.Class,
		&invitation.Section,
		&invitation.MissingUses,
		&invitation.CreateAt,
		&invitation.ServerId,
		&invitation.RoleId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invitation ID is invalid"})
		return
	}
	tokenResponse, err := ExchangeCodeForToken(discordCode.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Auth to Discord"})
		return
	}
	fmt.Println("Token Response")
	fmt.Println(tokenResponse.AccessToken)

	userId, err := getDiscordUserId(tokenResponse.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Auth to Discord"})
		return
	}
	fmt.Println("User id ")
	fmt.Println(userId)

	err = assignRoleToMember(invitation.ServerId, userId, invitation.RoleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot assign you a role, sorry."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Role %s assigned, to user %s \n", invitation.RoleId, userId),
		"role":    invitation.RoleId,
		"class":   invitation.Class,
		"section": invitation.Section,
	})

}

func ExchangeCodeForToken(code string) (*tokenResponse, error) {
	clientId := os.Getenv("DISCORD_CLIENT_ID")
	clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")

	// Create body
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:3000/#/assign-role/callback")

	// Creating requests
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, DISCORD_API_GET_TOKEN, strings.NewReader(data.Encode()))
	// Setting Authorization header
	r.SetBasicAuth(clientId, clientSecret)
	// Setting headers
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Executing requests
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Fetching BODY
	var tokenResponse tokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

func getDiscordUserId(token string) (string, error) {

	// Creating requests
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://discord.com/api/v10/users/@me", nil)
	r.Header.Add("Authorization", "Bearer "+token)

	// Executing requests
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Fetching BODY
	var discordUser discordUser

	if err := json.NewDecoder(resp.Body).Decode(&discordUser); err != nil {
		return "", err
	}

	return discordUser.ID, nil
}

func assignRoleToMember(serverId, userId, roleId string) error {

	uri := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s/roles/%s",
		serverId, userId, roleId)
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	fmt.Println("BOT TOKEN")
	fmt.Println(botToken)

	fmt.Println(uri)

	// Creating requests
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPut, uri, nil)
	r.Header.Add("Authorization", "Bot "+botToken)
	r.Header.Add("User-Agent", "IAmHumanBot (github.com/DanielRasho/IAmHuman-Bot, 1.0)")

	// Executing requests
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Println(resp.Status)
		return fmt.Errorf(resp.Status)
	}

	return nil
}
