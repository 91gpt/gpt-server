/**
 * @Author Nil
 * @Description api/image/generate.go
 * @Date 2023/4/10 20:43
 **/

package image

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ha5ky/hu5ky-bot/api"
	"github.com/ha5ky/hu5ky-bot/model"
	"github.com/ha5ky/hu5ky-bot/pkg/config"
	boterrors "github.com/ha5ky/hu5ky-bot/pkg/errors"
	"github.com/ha5ky/hu5ky-bot/pkg/logger"
	"github.com/sashabaranov/go-openai"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strconv"
)

type ImageResp struct {
	Images []string `json:"images"`
}

func Generate(ctx *gin.Context) {
	prompt := ctx.Query("prompt")
	openaiClient := openai.NewClient(config.SysCache.GPT.OpenaiAPIKey)
	var (
		ret       ImageResp
		agent     model.Agent
		user      model.User
		imageResp openai.ImageResponse
		uid       int
		err       error
	)
	if uid, err = strconv.Atoi(ctx.Query("user_id")); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
		return
	}
	userId := uint(uid)
	c := model.NewController()
	if user, err = c.UserModel(&model.User{}).Get(&model.UserQuery{
		ID: &userId,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
		return
	}
	if agent, err = c.AgentModel(&model.Agent{}).Get(&model.AgentQuery{
		ID: &user.AgentId,
	}); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
		return
	}
	req := openai.ImageRequest{
		Prompt: prompt,
		N:      agent.ImageN,
		Size:   agent.ImageSize,
	}

	ret.Images = make([]string, 0)
	if imageResp, err = openaiClient.CreateImage(ctx, req); err != nil {
		logger.Error(err)
		api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
		return
	}
	for i, item := range imageResp.Data {
		var (
			imageBytes []byte
			respTmp    = new(http.Response)
		)
		if respTmp, err = http.Get(item.URL); err != nil {
			logger.Error(err)
			api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
			return
		}
		if imageBytes, err = io.ReadAll(respTmp.Body); err != nil {
			logger.Error(err)
			api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
			return
		}
		ret.Images = append(ret.Images, base64.StdEncoding.EncodeToString(imageBytes))
		fileName := prompt + strconv.Itoa(i)
		if len(prompt) > 12 {
			fileName = prompt[:12] + strconv.Itoa(i)
		}
		if err = os.WriteFile(path.Join(config.SysCache.ServerConfig.Storage, "./storage", fileName+".png"), imageBytes, fs.ModePerm); err != nil {
			logger.Error(err)
			api.ErrorRender(ctx, http.StatusBadRequest, err, boterrors.InternalError)
			return
		}
	}
	api.OK(ctx, ret, 1)
}
