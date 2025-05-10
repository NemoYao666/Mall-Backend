package v1

import (
	"github.com/gin-gonic/gin"

	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

// InitSkillProductHandler 初始化秒杀商品信息
func InitSkillProductHandler(ctx *gin.Context) {
	var (
		req  types.ListSkillProductReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetSkillProductSrv()
	resp, err = l.InitSkillGoods(ctx.Request.Context())
	if err != nil {
		return
	}

}

// ListSkillProductHandler 初始化秒杀商品信息
func ListSkillProductHandler(ctx *gin.Context) {
	var (
		req  types.ListSkillProductReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetSkillProductSrv()
	resp, err = l.ListSkillGoods(ctx.Request.Context())
	if err != nil {
		return
	}

}

// GetSkillProductHandler 获取秒杀商品的详情
func GetSkillProductHandler(ctx *gin.Context) {
	var (
		req  types.GetSkillProductReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetSkillProductSrv()
	resp, err = l.GetSkillGoods(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func SkillProductHandler(ctx *gin.Context) {
	var (
		req  types.SkillProductReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetSkillProductSrv()
	resp, err = l.SkillProduct(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
