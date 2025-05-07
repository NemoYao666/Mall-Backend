package v1

import (
	"errors"
	"github.com/gin-gonic/gin"

	"gin-mall-backend/consts"
	"gin-mall-backend/pkg/e"
	"gin-mall-backend/service"
	"gin-mall-backend/types"
)

func UserRegisterHandler(ctx *gin.Context) {
	var (
		req  types.UserRegisterReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	// 参数校验
	if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
		err = errors.New("key长度错误,必须是6位数")
		return
	}

	var l = service.GetUserSrv()
	resp, err = l.UserRegister(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

// UserLoginHandler 用户登陆接口
func UserLoginHandler(ctx *gin.Context) {
	var (
		req  types.UserServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	var l = service.GetUserSrv()
	resp, err = l.UserLogin(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UserUpdateHandler(ctx *gin.Context) {
	var (
		req  types.UserInfoUpdateReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetUserSrv()
	resp, err = l.UserInfoUpdate(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func ShowUserInfoHandler(ctx *gin.Context) {
	var (
		req  types.UserInfoShowReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetUserSrv()
	resp, err = l.UserInfoShow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UploadAvatarHandler(ctx *gin.Context) {
	var (
		req  types.UserServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	file, fileHeader, _ := ctx.Request.FormFile("file")
	if fileHeader == nil {
		err = errors.New(e.GetMsg(e.ErrorUploadFile))
		return
	}

	l := service.GetUserSrv()
	resp, err = l.UserAvatarUpload(ctx.Request.Context(), file, fileHeader.Filename, &req)
	if err != nil {
		return
	}

}

func SendEmailHandler(ctx *gin.Context) {
	var (
		req  types.SendEmailServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetUserSrv()
	resp, err = l.SendEmail(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UserFollowingHandler(ctx *gin.Context) {
	var (
		req  types.UserFollowingReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetUserSrv()
	resp, err = l.UserFollow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UserUnFollowingHandler(ctx *gin.Context) {
	var (
		req  types.UserUnFollowingReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	l := service.GetUserSrv()
	resp, err = l.UserUnFollow(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UserFollowingListHandler(ctx *gin.Context) {
	var (
		req  types.UserFollowingListReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	var l = service.GetUserSrv()
	resp, err = l.UserFollowingList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}

func UserFollowerListHandler(ctx *gin.Context) {
	var (
		req  types.UserFollowerListReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	var l = service.GetUserSrv()
	resp, err = l.UserFollowerList(ctx.Request.Context(), &req)
	if err != nil {
		return
	}
}

func ValidEmailHandler(ctx *gin.Context) {
	var (
		req  types.ValidEmailServiceReq
		resp interface{}
		err  error
	)
	defer func() {
		StandardResponse(ctx, err, resp)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		return
	}

	var l = service.GetUserSrv()
	resp, err = l.Valid(ctx.Request.Context(), &req)
	if err != nil {
		return
	}

}
