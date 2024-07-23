package handler

import (
	"github.com/gin-gonic/gin"
	controler "github.com/online_marketplace/internal/controller"
	"github.com/online_marketplace/internal/registry"
)

func registerProductRouter(appCtx *registry.ServiceContext, api *gin.RouterGroup) {
	var (
		path = "/products"
	)
	user := api.Group(path)
	{
		user.POST("", controler.NewProductController(appCtx).Create())
		user.GET("", controler.NewProductController(appCtx).List())
	}
}
