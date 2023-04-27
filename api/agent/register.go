/**
 * @Author Nil
 * @Description api/agent/register.go
 * @Date 2023/4/24 20:03
 **/

package agent

import "github.com/ha5ky/hu5ky-bot/router/schema"

func init() {
	scheme := schema.NewSchemeBuilder().Register()

	scheme.POST("/agent", Create)
	scheme.GET("/agents", List)
	scheme.PUT("/agent", Update)
}
