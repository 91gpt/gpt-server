/**
 * @Author Nil
 * @Description api/user/register.go
 * @Date 2023/4/21 10:53
 **/

package user

import "github.com/ha5ky/hu5ky-bot/router/schema"

func init() {
	scheme := schema.NewSchemeBuilder().Register()

	scheme.GET("/user", Login)
	scheme.POST("/user", Generate)
}
