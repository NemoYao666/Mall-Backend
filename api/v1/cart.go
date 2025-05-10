package v1

import (
	"gin-mall-backend/consts"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

func CreateCartHandler(ctx *gin.Context) {
	var (
		req  types.CartCreateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetCartSrv()
	resp, err = l.CartCreate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// ListCartHandler 购物车详细信息
func ListCartHandler(ctx *gin.Context) {
	var (
		req  types.CartListReq
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

	l := service.GetCartSrv()
	resp, err = l.CartList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// UpdateCartHandler 修改购物车信息
func UpdateCartHandler(ctx *gin.Context) {
	var (
		req  types.UpdateCartServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetCartSrv()
	resp, err = l.CartUpdate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// DeleteCartHandler 删除购物车
func DeleteCartHandler(ctx *gin.Context) {
	var (
		req  types.CartDeleteReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetCartSrv()
	resp, err = l.CartDelete(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
