package controler

import (
	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/server/http/response"
	"github.com/online_marketplace/internal/registry"
	"github.com/online_marketplace/internal/service"
	"github.com/online_marketplace/internal/types/request"
)

type UserController struct {
	reg     *registry.ServiceContext
	userSvc service.UserService
}

func NewUserController(reg *registry.ServiceContext) *UserController {
	return &UserController{
		reg:     reg,
		userSvc: service.NewUserService(reg),
	}
}

func (uc *UserController) Register() func(*gin.Context) {
	return func(c *gin.Context) {
		var req request.RegisterReq
		ctx := c.Request.Context()

		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}

		if err := req.Validate(ctx); err != nil {
			panic(err)

		}

		resp, err := uc.userSvc.Register(ctx, req)
		if err != nil {
			panic(err)
		}

		response.SuccessResponse(c, resp, nil)
	}

}
