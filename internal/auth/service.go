package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type service struct {
	baseUrl string
	client  *resty.Client
}

type Interface interface {
	GetAll(ctx *gin.Context)
}

func NewAuthService(baseUrl string) Interface {
	return &service{
		baseUrl: baseUrl,
		client:  resty.New(),
	}
}
