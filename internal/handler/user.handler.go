package handler

import (
	"github.com/gin-gonic/gin"
	controler "github.com/online_marketplace/internal/controller"
	"github.com/online_marketplace/internal/registry"
)

func registerUserRouter(appCtx *registry.ServiceContext, api *gin.RouterGroup) {
	var (
		path = "/users"
	)
	user := api.Group(path)
	{
		user.POST("/register", controler.NewUserController(appCtx).Register())
	}
}
