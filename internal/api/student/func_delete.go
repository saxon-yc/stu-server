package stuapi

import (
	"net/http"
	"strconv"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 删除学生列表
// @Schemes
// @Description 删除学生列表
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param id query string true "删除学生"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/student/:id [delete]
func (h *stuHandle) DeleteStudentAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, _ := strconv.Atoi(id)
		err := h.websvc.DeleteStudent(uint32(i))
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "删除成功"})
	}
}
