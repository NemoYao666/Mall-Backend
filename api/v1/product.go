package v1

import (
	"errors"
	"github.com/gin-gonic/gin"

	"gin-mall-backend/consts"
	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

// CreateProductHandler 创建商品
func CreateProductHandler(ctx *gin.Context) {
	var (
		req  types.ProductCreateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	form, _ := ctx.MultipartForm()
	files := form.File["image"]
	l := service.GetProductSrv()
	resp, err = l.ProductCreate(ctx.Request.Context(), files, &req)
	if err != nil {
		return
	}

}

// ListProductsHandler 商品列表
func ListProductsHandler(ctx *gin.Context) {
	var (
		req  types.ProductListReq
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
		req.PageSize = consts.BaseProductPageSize
	}

	l := service.GetProductSrv()
	resp, err = l.ProductList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// ShowProductHandler 商品详情
func ShowProductHandler(ctx *gin.Context) {
	var (
		req  types.ProductShowReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetProductSrv()
	resp, err = l.ProductShow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// DeleteProductHandler 删除商品
func DeleteProductHandler(ctx *gin.Context) {
	var (
		req  types.ProductDeleteReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetProductSrv()
	resp, err = l.ProductDelete(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// UpdateProductHandler 更新商品
func UpdateProductHandler(ctx *gin.Context) {
	var (
		req  types.ProductUpdateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetProductSrv()
	resp, err = l.ProductUpdate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// SearchProductsHandler 搜索商品
func SearchProductsHandler(ctx *gin.Context) {
	var (
		req  types.ProductSearchReq
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

	l := service.GetProductSrv()
	resp, err = l.ProductSearch(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func ListProductImgHandler(ctx *gin.Context) {
	var (
		req  types.ListProductImgReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	if req.ID == 0 {
		err = errors.New("参数错误,id不能为空")
		return
	}

	l := service.GetProductSrv()
	resp, err = l.ProductImgList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
