/**
 * @Author Nil
 * @Description api/user/login.go
 * @Date 2023/4/21 10:54
 **/

package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
)

type CallBack struct {
	UserId uint `json:"user_id"`
}

func Login(ctx *gin.Context) {
	account := ctx.Query("account")
	pwd := ctx.Query("pwd")
	var (
		uid uint
		err error
		ok  bool
	)
	c := model.NewController()
	if uid, ok, err = c.UserModel(&model.User{}).Check(&model.UserQuery{
		Account: &account,
		Pwd:     &pwd,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	if !ok {
		err = errors.New(boterrors.InvalidAuth)
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidAuth)
		return
	}
	q := &model.UserQuery{ID: &uid}
	if err = c.UserModel(&model.User{}).Activate(q); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	ret := CallBack{
		UserId: uid,
	}
	api.OK(ctx, ret, 0)
}
