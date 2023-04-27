/**
 * @Author Nil
 * @Description api/apikey/add.go
 * @Date 2023/4/21 15:47
 **/

package apikey

import (
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"net/http"
)

type AddReq struct {
	ApiKey string `json:"api_key" form:"api_key" binding:"required"`
}

func Add(ctx *gin.Context) {
	var (
		req AddReq
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		logger.Info("Invalid api-key")
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidApiKey)
		return
	}

}
