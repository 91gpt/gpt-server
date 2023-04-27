/**
 * @Author Nil
 * @Description api/history/register.go
 * @Date 2023/4/24 16:33
 **/

package chat

import (
	"github.com/ha5ky/hu5ky-bot/router/schema"
)

func init() {
	scheme := schema.NewSchemeBuilder().Register().Group("/history")

	scheme.GET("/chat/:id", Get)
	scheme.POST("/chat", Create)
	scheme.GET("/chats/:user_id", List)
}
