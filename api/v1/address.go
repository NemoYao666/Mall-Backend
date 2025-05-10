package v1

import (
	"gin-mall-backend/consts"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

// CreateAddressHandler 新增收货地址
func CreateAddressHandler(ctx *gin.Context) {
	var (
		req  types.AddressCreateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetAddressSrv()
	resp, err = l.AddressCreate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// ShowAddressHandler 展示某个收货地址
func ShowAddressHandler(ctx *gin.Context) {
	var (
		req  types.AddressGetReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetAddressSrv()
	resp, err = l.AddressShow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// ListAddressHandler 展示收货地址
func ListAddressHandler(ctx *gin.Context) {
	var (
		req  types.AddressListReq
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

	l := service.GetAddressSrv()
	resp, err = l.AddressList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// UpdateAddressHandler 修改收货地址
func UpdateAddressHandler(ctx *gin.Context) {
	var (
		req  types.AddressServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetAddressSrv()
	resp, err = l.AddressUpdate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// DeleteAddressHandler 删除收获地址
func DeleteAddressHandler(ctx *gin.Context) {
	var (
		req  types.AddressDeleteReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetAddressSrv()
	resp, err = l.AddressDelete(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
