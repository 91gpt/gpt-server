/**
 * @Author Nil
 * @Description api/apikey/check.go
 * @Date 2023/4/19 20:08
 **/

package apikey

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func Check(ctx *gin.Context) {
	apiKey := ctx.Param("apiKey")
	openaiClient := openai.NewClient(apiKey)
	var (
		modelsList openai.ModelsList
		err        error
	)
	if modelsList, err = openaiClient.ListModels(ctx); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidParams)
		return
	}
	if len(modelsList.Models) == 0 {
		err = errors.New(boterrors.InvalidApiKey)
		logger.Info("Invalid api-key")
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InvalidApiKey)
		return
	}
	api.OK(ctx, nil, 0)
}
