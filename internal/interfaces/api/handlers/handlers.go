package handlers

import (
	"net/http"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/security"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/interfaces/api/middleware"
	"github.com/dh-atha/EmployeeAbsenceKNTest/pkg/config"
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Data:    data,
	}
}

type Handlers struct {
	jwtService        security.JWTService
	authHandler       *AuthHandler
	membershipHandler *MembershipHandler
	contactHandler    *ContactHandler
}

type HandlersRequirements struct {
	JwtService        security.JWTService
	AuthHandler       *AuthHandler
	MembershipHandler *MembershipHandler
	ContactHandler    *ContactHandler
}

func NewHandlers(req HandlersRequirements) *Handlers {
	return &Handlers{
		jwtService:        req.JwtService,
		authHandler:       req.AuthHandler,
		membershipHandler: req.MembershipHandler,
		contactHandler:    req.ContactHandler,
	}
}

func (a *Handlers) CreateServer(address string) (*http.Server, error) {
	gin.SetMode(config.Configuration.Server.Mode)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/ping", a.checkConnectivity)

	r.POST("/membership", a.membershipHandler.CreateMembership)
	r.POST("/contact", a.contactHandler.CreateContact)
	r.POST("/login", a.authHandler.Login)
	r.Use(middleware.AuthMiddleware(a.jwtService))

	r.GET("/membership", a.membershipHandler.GetAllMemberships)
	r.PUT("/contact/:id", a.contactHandler.UpdateContact)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return server, nil
}

func (a *Handlers) checkConnectivity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
