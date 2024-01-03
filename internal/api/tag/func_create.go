package tagapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 新建标签
// @Schemes
// @Description 新建标签
// @Tags 标签管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.CreateTagRequest{} true "查询标签接口参数"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/tag [post]
func (h *tagHandle) CreateTagAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateTagRequest
		c.ShouldBindJSON(&req)
		err := h.websvc.CreateTag(req)

		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "创建成功"})
	}
}
