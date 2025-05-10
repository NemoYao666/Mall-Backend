package v1

import (
	"gin-mall-backend/service"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/types"
)

func OrderPaymentHandler(ctx *gin.Context) {
	var (
		req  types.PaymentDownReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetPaymentSrv()
	resp, err = l.PayDown(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
