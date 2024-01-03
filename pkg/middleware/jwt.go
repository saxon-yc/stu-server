package middleware

import (
	"net/http"
	"strings"
	"student-server/internal/model"
	errorcode "student-server/pkg/errors"
	"student-server/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取到请求头中的token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, &model.BaseResponse{
				Code: errorcode.INVALID_PASSWORD_FAILED_CODE,
				Msg:  "访问失败: 请重新登录!",
				Data: nil,
			})
			// 中断本次请求
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, &model.BaseResponse{
				Code: errorcode.INVALID_PASSWORD_FAILED_CODE,
				Msg:  "访问失败: 无效的token, 请重新登录!",
				Data: nil,
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, &model.BaseResponse{
				Code: errorcode.INVALID_PASSWORD_FAILED_CODE,
				Msg:  "访问失败: 无效的token, 请重新登录!",
				Data: nil,
			})
			c.Abort()
			return
		}
		// 将当前请求的account信息保存到请求的上下文c上
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Set("username", mc.UserName)
		c.Next()
	}
}
