/**
 * @Author Nil
 * @Description api/agent/list.go
 * @Date 2023/4/26 10:24
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

func List(ctx *gin.Context) {
	var (
		agents = make([]*model.Agent, 0)
		total  int64
		err    error
	)
	c := model.NewController()
	if agents, total, err = c.AgentModel(&model.Agent{}).List(&model.AgentQuery{}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	api.OK(ctx, agents, int(total))
}
