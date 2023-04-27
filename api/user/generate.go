/**
 * @Author Nil
 * @Description api/user/generate.go
 * @Date 2023/4/21 13:40
 **/

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
)

type GenerateReq struct {
	MainUserId uint   `json:"main_user_id" form:"main_user_id" binding:"required"`
	AgentId    uint   `json:"agent_id" form:"agent_id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
}

func Generate(ctx *gin.Context) {
	var (
		req GenerateReq
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		logger.Info(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}

	var (
		agent model.Agent
	)

	c := model.NewController()
	_ = c.Begin()
	if agent, err = c.AgentModel(&model.Agent{}).Get(&model.AgentQuery{
		ID: &req.AgentId,
	}); err != nil {
		c.Rollback()
		logger.Info(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.DatabaseQueryError)
		return
	}
	if err = c.UserModel(&model.User{}).Generate(req.MainUserId, &agent); err != nil {
		c.Rollback()
		logger.Info(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.DatabaseQueryError)
		return
	}
	c.Commit()
	api.OK(ctx, nil, 0)
}
