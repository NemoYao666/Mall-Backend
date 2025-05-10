package v1

import (
	"gin-mall-backend/consts"
	"gin-mall-backend/service"
	"gin-mall-backend/types"

	"github.com/gin-gonic/gin"
)

// CreateFavoriteHandler 创建收藏
func CreateFavoriteHandler(ctx *gin.Context) {
	var (
		req  types.FavoriteCreateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetFavoriteSrv()
	resp, err = l.FavoriteCreate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// ListFavoritesHandler 收藏夹详情接口
func ListFavoritesHandler(ctx *gin.Context) {
	var (
		req  types.FavoritesServiceReq
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

	l := service.GetFavoriteSrv()
	resp, err = l.FavoriteList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// DeleteFavoriteHandler 删除收藏夹
func DeleteFavoriteHandler(ctx *gin.Context) {
	var (
		req  types.FavoriteDeleteReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetFavoriteSrv()
	resp, err = l.FavoriteDelete(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
