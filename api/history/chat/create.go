/**
 * @Author Nil
 * @Description api/history/chat/create.go
 * @Date 2023/4/24 17:05
 **/

package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
)

type CreateReq struct {
	UserId uint   `json:"user_id" form:"user_id" binding:"required"`
	Title  string `json:"title" form:"title" binding:"required"`
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

	h := &model.HistoricMsg{
		UserId: req.UserId,
		Title:  req.Title,
	}
	c := model.NewController()
	if err = c.HistoricMsgModel(h).Save(); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	api.OK(ctx, nil, 0)
}
