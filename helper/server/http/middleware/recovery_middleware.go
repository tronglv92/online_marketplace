package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/errors"
	"github.com/online_marketplace/helper/server/core"
	"github.com/online_marketplace/helper/server/http/response"
)

type RecoveryMiddleware struct {
	EnvMode string
}

func NewRecoveryMiddleware(mode string) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		EnvMode: mode,
	}
}

func (m *RecoveryMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// if err := recover(); err != nil {
			// 	c.Header("Content-Type", "application/json")

			// 	if appErr, ok := err.(*common.AppError); ok {
			// 		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			// 		panic(err)
			// 		return
			// 	}

			// 	appErr := common.ErrInternal(err.(error))
			// 	c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			// 	panic(err)
			// 	return
			// }
			if result := recover(); result != nil {

				err := errors.ToError(result)
				e := errors.From(err)

				response.ErrorResponse(c, errors.Newf(
					e.GetCode(),
					e.GetReason(),
					errors.WithStack(core.SprintStack()),
				))

				c.Abort()
			}
		}()

		c.Next()
	}
}
