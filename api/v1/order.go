package v1

import (
	"gin-mall-backend/consts"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

func CreateOrderHandler(ctx *gin.Context) {
	var (
		req  types.OrderCreateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetOrderSrv()
	resp, err = l.OrderCreate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}
	return

}

func ListOrdersHandler(ctx *gin.Context) {
	var (
		req  types.OrderListReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	if req.PageSize == 0 {
		req.PageSize = consts.BasePageSize
	}

	l := service.GetOrderSrv()
	resp, err = l.OrderList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}
	return

}

// ShowOrderHandler 订单详情
func ShowOrderHandler(ctx *gin.Context) {
	var (
		req  types.OrderShowReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetOrderSrv()
	resp, err = l.OrderShow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}
	return

}

func DeleteOrderHandler(ctx *gin.Context) {
	var (
		req  types.OrderDeleteReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetOrderSrv()
	resp, err = l.OrderDelete(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
