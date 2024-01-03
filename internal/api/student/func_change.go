package stuapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 修改学生
// @Schemes
// @Description 修改学生
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param params body model.ChangeStudentRequest{} true "修改学生信息"
// @Success 200 {object} model.BaseResponse{} "desc"
// @Router /v2/student [put]
func (h *stuHandle) ChangeStudentAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ChangeStudentRequest
		c.ShouldBindJSON(&req)
		err := h.websvc.ChangeStudent(req)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{Code: 1, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.BaseResponse{Code: 0, Msg: "修改学生信息成功"})
	}
}
