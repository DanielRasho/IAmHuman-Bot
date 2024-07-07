package server

import (
	"backend/internal/database"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port int

	db database.Service

	engine *gin.Engine
}

type userCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type invitation struct {
	Id            string    `json:"id"`
	Class         string    `json:"class"`
	Section       int       `json:"section"`
	MissingUses   int       `json:"missing_uses"`
	CreateAt      time.Time `json:"created_at"`
	ServerId      string    `json:"server_id"`
	RoleId        string    `json:"role_id"`
	InvitationURL string    `json:"invitation_url"`
}

type discordCode struct {
	Code         string `json:"code"`
	InvitationId string `json:"invitation_id"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type discordAvatarDecoration struct {
	SkuID string `json:"sku_id"`
	Asset string `json:"asset"`
}

type discordUser struct {
	ID               string                  `json:"id"`
	Username         string                  `json:"username"`
	Discriminator    string                  `json:"discriminator"`
	Avatar           string                  `json:"avatar"`
	Verified         bool                    `json:"verified"`
	Email            string                  `json:"email"`
	Flags            int                     `json:"flags"`
	Banner           string                  `json:"banner"`
	AccentColor      int                     `json:"accent_color"`
	PremiumType      int                     `json:"premium_type"`
	PublicFlags      int                     `json:"public_flags"`
	AvatarDecoration discordAvatarDecoration `json:"avatar_decoration_data"`
}
