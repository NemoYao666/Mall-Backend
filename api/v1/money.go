package v1

import (
	"gin-mall-backend/service"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/types"
)

func ShowMoneyHandler(ctx *gin.Context) {
	var (
		req  types.MoneyShowReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetMoneySrv()
	resp, err = l.MoneyShow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
