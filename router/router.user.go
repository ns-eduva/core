package router

import (
	"eduva/core/di"
	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine, dis *di.AppDependencies) {
	v1 := r.Group(pathApiV1)
	v1User := v1.Group("/user")
	v1User.GET("/", dis.UserController.GetAll)
}
