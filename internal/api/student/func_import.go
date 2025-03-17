package stuapi

import (
	"fmt"
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 导入学生列表
// @Schemes
// @Description 导入学生列表
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.ImportStudentsAPI true "关键字查询"
// @Success 200 {object} string "desc"
// @Router /v2/student/import [post]
func (h *stuHandle) ImportStudentsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		total, err := h.websvc.ImportStudentsAPI(file)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.BaseResponse{
			Code: 0,
			Msg:  fmt.Sprintf("导入成功，共导入 %d 条数据", total),
		})
	}
}
