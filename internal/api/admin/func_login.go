package adminapi

import (
	"net/http"
	"strings"
	"student-server/internal/model"
	errorcode "student-server/pkg/errors"
	"student-server/pkg/util"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 登录接口
// @Schemes
// @Description 登录接口
// @Tags 用户管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.LoginRequest{} true "登录接口参数"
// @Success 200 {object} model.BaseResponse{data=model.UserResponse} "desc"
// @Router /v1/login [post]
func (h *handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.LoginRequest)
		c.ShouldBindJSON(req)
		params := model.LoginRequest{
			UserName:        req.UserName,
			Password:        util.GeneratePassword(req.Password),
			IsCreateAccount: req.IsCreateAccount,
		}

		user, err := h.websvc.LoginWeb(c, params)

		if err != nil {
			// 类型断言: 将变量 err 转换为 ErrorCodeString 类型的变量。
			ec, ok := err.(*errorcode.ErrorCodeString)
			if !ok {
				ec = errorcode.New(errorcode.UNKOWN_ERROR_CODE, "Unkown", err.Error())
			} else {
				if strings.Contains(ec.Error(), "内部命令执行错误") {
					ec.SetErr("当前版本不支持该操作")
				}
			}
			c.JSON(http.StatusOK, &model.BaseResponse{
				Code: ec.Code(),
				Msg:  ec.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &model.BaseResponse{
			Code: 0,
			Msg:  "登录成功",
			Data: &model.UserResponse{
				Token:    user.Token,
				UserName: user.UserName,
				Menus:    user.Menus,
			},
		})
	}
}
