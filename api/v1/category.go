package v1

import (
	"gin-mall-backend/service"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/types"
)

func ListCategoryHandler(ctx *gin.Context) {
	var (
		req  types.ListCategoryReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetCategorySrv()
	resp, err = l.CategoryList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
