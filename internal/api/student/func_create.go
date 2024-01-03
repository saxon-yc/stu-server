package stuapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 新建学生
// @Schemes
// @Description 新建学生
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.CreateStudentRequest{} true "新建学生参数"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/student [post]
func (h *stuHandle) CreateStudentAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateStudentRequest
		c.ShouldBindJSON(&req)
		err := h.websvc.CreateStudent(req)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "新建学生成功"})
	}
}
