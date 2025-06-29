package usercontroller

import (
	"eduva/core/internal/auth"
	"eduva/core/internal/crm"
	"github.com/gin-gonic/gin"
)

type controller struct {
	crmService  crm.Interface
	authService auth.Interface
}

type Interface interface {
	GetAll(ctx *gin.Context)
}

func NewUserController(crm crm.Interface, auth auth.Interface) Interface {
	return &controller{
		crmService:  crm,
		authService: auth,
	}
}
