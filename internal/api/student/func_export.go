package stuapi

import (
	"net/http"
	"student-server/internal/model"
	"student-server/pkg/excelizelib"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 导出学生列表
// @Schemes
// @Description 导出学生列表
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.ExportStudentRequest true "关键字查询"
// @Success 200 {object} string "desc"
// @Router /v2/student/export [post]
func (h *stuHandle) ExportStudentsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ExportStudentRequest
		c.ShouldBindJSON(&req)

		dataKey, data, err := h.websvc.ExportStudent(req)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		excel := excelizelib.NewMyExcel("学生信息", "学生信息")
		excel.ExportToWeb(dataKey, data, c)

	}
}
