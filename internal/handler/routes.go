package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/server/http/middleware"
	"github.com/online_marketplace/internal/config"
	"github.com/online_marketplace/internal/registry"
)

const (
	BasePrefix = "/online-marketplace-svc"
	RestPrefix = BasePrefix + "/api/v1"
)

func MustSetup(c config.ServerConfig) {
	Initialize()

}
func Initialize() {
	for _, v := range Providers() {
		v.Register()
	}
}

func NewRouter(appCtx *registry.ServiceContext) *gin.Engine {
	MustSetup(appCtx.Config.Server)

	r := gin.Default()

	r.Use(middleware.NewRecoveryMiddleware(appCtx.Config.Server.Env).Handle())
	api := r.Group(RestPrefix)
	{
		registerUserRouter(appCtx, api)
	}

	return r
}
