/**
 * @Author Nil
 * @Description api/history/chat/list.go
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
	"strconv"
)

func List(ctx *gin.Context) {
	var (
		hs      = make([]*model.HistoricMsg, 0)
		uid     int
		total   int64
		uintUId uint
		err     error
	)
	userId := ctx.Param("user_id")
	if uid, err = strconv.Atoi(userId); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	uintUId = uint(uid)
	c := model.NewController()
	if hs, total, err = c.HistoricMsgModel(&model.HistoricMsg{}).List(&model.HistoricMsgQuery{
		UserId:  &uintUId,
		PreLoad: true,
		Desc:    true,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	api.OK(ctx, hs, int(total))
}
