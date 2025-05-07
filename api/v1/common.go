package v1

import (
	"encoding/json"
	"fmt"
	"gin-mall-backend/consts"
	"gin-mall-backend/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	conf "gin-mall-backend/config"
	"gin-mall-backend/pkg/e"
	"gin-mall-backend/pkg/utils/ctl"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.TrackedErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", fieldError.Tag))
			return ctl.RespError(ctx, err, fmt.Sprintf("%s%s", field, tag))
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(ctx, err, "JSON类型不匹配")
	}

	return ctl.RespError(ctx, err, err.Error(), e.ERROR)
}

func StandardResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(consts.StatusOK, ctl.RespSuccess(ctx, data))
		return
	}

	log.LogrusObj.Infoln(err)
	ctx.JSON(consts.ServerError, ErrorResponse(ctx, err))
	return
}
