/**
 * @Author Nil
 * @Description api/apikey/register.go
 * @Date 2023/4/19 20:09
 **/

package apikey

import "github.com/ha5ky/hu5ky-bot/router/schema"

func init() {
	scheme := schema.NewSchemeBuilder().Register()

	scheme.GET("/api-key/:apiKey", Check)
}
