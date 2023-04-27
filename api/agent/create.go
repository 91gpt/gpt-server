/**
 * @Author Nil
 * @Description api/agent/create.go
 * @Date 2023/4/24 20:08
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

type CreateReq struct {
	ImageN           int    `json:"image_n" form:"image_n"`
	ImageSize        string `json:"image_size" form:"image_size"`
	PID              uint   `json:"pid" form:"pid" binding:"required"`
	SubAccountNumber int    `json:"sub_account_number" form:"sub_account_number" binding:"required"`
	Daily            int    `json:"daily" form:"daily" binding:"required"`
}

func Create(ctx *gin.Context) {
	var (
		req CreateReq
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}

	var (
		pAgent model.Agent
		c      = model.NewController()
	)
	if pAgent, err = c.AgentModel(&model.Agent{}).Get(&model.AgentQuery{
		ID: &req.PID,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}

	agent := &model.Agent{
		PID:              req.PID,
		ImageN:           req.ImageN,
		ImageSize:        req.ImageSize,
		SubAccountNumber: req.SubAccountNumber,
		Daily:            req.Daily,
		Level:            pAgent.Level + 1,
	}
	if err = c.AgentModel(agent).Save(); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	api.OK(ctx, nil, 0)
}
