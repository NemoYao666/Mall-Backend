package v1

import (
	"gin-mall-backend/service"

	"github.com/gin-gonic/gin"

	"gin-mall-backend/types"
)

func ListCarouselsHandler(ctx *gin.Context) {
	var (
		req  types.ListCarouselReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetCarouselSrv()
	resp, err = l.ListCarousel(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
