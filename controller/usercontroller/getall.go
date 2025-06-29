package usercontroller

import "github.com/gin-gonic/gin"

func (c *controller) GetAll(ctx *gin.Context) {
	c.authService.GetAll(ctx)
}
