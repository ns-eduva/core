package router

import (
	"eduva/core/di"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const pathApiV1 = "api/v1"

func New(r *gin.Engine, dis *di.AppDependencies) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // TODO : a déplacer quand le router sera fait

	RouterUser(r, dis)
}
