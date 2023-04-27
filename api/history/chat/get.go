/**
 * @Author Nil
 * @Description api/history/chat.go
 * @Date 2023/4/24 16:36
 **/

package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
	"strconv"
)

func Get(ctx *gin.Context) {
	var (
		id     = ctx.Param("id")
		uid    int
		uintId uint
		err    error

		h model.HistoricMsg
	)
	if uid, err = strconv.Atoi(id); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	uintId = uint(uid)
	c := model.NewController()
	if h, err = c.HistoricMsgModel(&model.HistoricMsg{}).Get(&model.HistoricMsgQuery{
		PreLoad: true,
		ID:      &uintId,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.DatabaseQueryError)
		return
	}
	api.OK(ctx, h, 1)
}
