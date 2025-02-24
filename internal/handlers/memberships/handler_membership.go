package memberships

import (
	"context"
	"github.com/gin-gonic/gin"
	"mySimpleFprum/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, request memberships.LoginRequest) (string, string, error)
}

type Handler struct {
	Engine        *gin.Engine
	membershipSvc membershipService
}

func NewHandler(engine *gin.Engine, membershipsSvc membershipService) *Handler {
	return &Handler{
		Engine:        engine,
		membershipSvc: membershipsSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	routes := h.Engine.Group("memberships")
	//routes.Use(middleware.AuthMiddleware)
	routes.POST("/signup", h.SignUp)
	routes.POST("/login", h.Login)
}
