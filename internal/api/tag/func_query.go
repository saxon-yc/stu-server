package tagapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 查询标签列表
// @Schemes
// @Description 查询标签列表
// @Tags 标签管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param search_word query string false "关键字查询"
// @Param limit query int false "页数"
// @Param offset query int false "页码"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Success 200 {object} model.BaseResponse{data=model.QueryTagResponse} "desc"
// @Router /v2/tag [get]
func (h *tagHandle) QueryTagsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.QueryTagRequest
		c.ShouldBindQuery(&req)
		result, err := h.websvc.QueryTag(req)

		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.BaseResponse{
			Code: 0,
			Data: result,
			Msg:  "success",
		})
	}
}
