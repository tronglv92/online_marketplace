package response

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/errors"
	"github.com/online_marketplace/helper/locale"
)

func setHttpResponse(ctx context.Context, result any, paging any, err errors.Error) any {
	dt := data{}
	var msgKey, msgResp = locale.SuccessMsg.Key, locale.SuccessMsg.Message
	if err != nil {

		msgKey, msgResp = fmt.Sprintf("%d", err.GetCode()), err.GetReason()
	}
	return responseHttp{
		Meta: metaResponse{
			TradeId: "123",
			Code:    msgKey,
			Message: msgResp,
		},
		Data: dt.SetData(result, paging),
	}
}

func parseError(ctx context.Context, err error) (int, errors.Error) {
	// span := trace.SpanFromContext(ctx)
	// defer span.End()

	e := errors.From(err)
	// if e.HasReport() {
	// 	logify.NewReport().Send(ctx, e)
	// }
	// if span.IsRecording() {
	// 	span.RecordError(err)
	// }
	return e.GetCode(), e
}

func ErrorResponse(c *gin.Context, err error) {
	ctx := c.Request.Context()
	status, e := parseError(ctx, err)
	c.JSON(status, setHttpResponse(ctx, nil, nil, e))
}
func SuccessResponse(c *gin.Context, result any, paging any) {
	c.JSON(http.StatusOK, setHttpResponse(c.Request.Context(), result, paging, nil))
	// httpx.OkJsonCtx(ctx, w, setHttpResponse(ctx, result, paging, nil))
	// return
}
