/**
 * @Author Nil
 * @Description api/agent/update.go
 * @Date 2023/4/26 10:30
 **/

package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
)

type UpdateReq struct {
	Id               uint   `json:"id" form:"id" binding:"required"`
	ImageN           int    `json:"image_n" form:"image_n"`
	ImageSize        string `json:"image_size" form:"image_size"`
	SubAccountNumber int    `json:"sub_account_number" form:"sub_account_number" binding:"required"`
	Daily            int    `json:"daily" form:"daily" binding:"required"`
}

func Update(ctx *gin.Context) {
	var (
		req UpdateReq
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}

	var (
		agent model.Agent
		c     = model.NewController()
	)
	if agent, err = c.AgentModel(&model.Agent{}).Get(&model.AgentQuery{
		ID: &req.Id,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}

	agentUpdate := &model.Agent{
		Model:            agent.Model,
		PID:              agent.PID,
		ImageN:           req.ImageN,
		ImageSize:        req.ImageSize,
		SubAccountNumber: req.SubAccountNumber,
		Daily:            req.Daily,
		Level:            agent.Level,
	}
	if err = c.AgentModel(agentUpdate).Save(); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	api.OK(ctx, nil, 0)
}
