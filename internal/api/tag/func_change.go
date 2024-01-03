package tagapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 修改标签
// @Schemes
// @Description 修改标签
// @Tags 标签管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.ChangeTagRequest{} true "查询标签接口参数"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/tag [patch]
func (h *tagHandle) ChangeTagAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ChangeTagRequest
		c.ShouldBindJSON(&req)
		err := h.websvc.ChangeTag(req)

		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "更新成功"})
	}
}
