package middleware

import (
	"apiserver_demos/demo09/handler"
	"apiserver_demos/demo09/pkg/errno"
	"apiserver_demos/demo09/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			log.Error("parse token occur err", err)
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
