package middlewares

import (
	"net/http"
	"time"

	"github.com/cindyyangcaixia/gin-example/pkg/app"
	"github.com/cindyyangcaixia/gin-example/pkg/e"
	"github.com/cindyyangcaixia/gin-example/pkg/utils"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		schema := c.Request.Header.Get("x-auth-schema")

		code := e.SUCCESS
		token := utils.GetToken(authorization)

		if token == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := utils.VerifyToken(token, schema)

			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = e.TOKEN_EXPIRED
			}
		}

		if code != e.SUCCESS {
			appGin := app.Gin{C: c}
			var data interface{}
			appGin.Response(&app.ResponseMessage{Status: http.StatusUnauthorized, Code: code}, data)
			c.Abort()
			return
		}
		c.Next()
	}
}
