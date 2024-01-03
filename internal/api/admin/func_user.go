package adminapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address interface{} `json:"address"`
}

// PingExample godoc
// @Summary 重名用户
// @Schemes
// @Description 查询用户是否已存在
// @Tags 用户管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param username query string true "用户姓名"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Failure 400 {object} model.BaseResponse{} "desc"
// @Router /v1/dupuser [get]
func (h *handler) QueryUserExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.DupuserRequest
		c.ShouldBindQuery(&req)
		_, err := h.websvc.QueryUser(req)
		if err == nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: "用户已存在"})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "success"})

	}
}

// PingExample godoc
// @Summary 用户信息
// @Schemes
// @Description 用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username body string true "账号名"
// @Success 200 {object} JSONResult{data=User} "desc"
// @Failure 400 {object} JSONResult{} "desc"
// @Router /v1/user [get]
func (h *handler) QueryUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "用户信息")
	}
}
