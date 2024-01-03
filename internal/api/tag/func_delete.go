package tagapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 删除标签列表
// @Schemes
// @Description 删除标签列表
// @Tags 标签管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param id query string true "删除标签"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/tag/:id [delete]
func (h *tagHandle) DeleteTagAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.DeleteTagRequest
		c.ShouldBindUri(&req)
		err := h.websvc.DeleteTag(req)

		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "删除成功"})
	}
}
