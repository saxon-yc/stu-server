package stuapi

import (
	"net/http"
	"student-server/internal/model"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary 查询学生列表
// @Schemes
// @Description 查询学生列表
// @Tags 学生管理
// @Accept application/json;charset=utf-8
// @Produce json
// @Param search_word query string false "关键字查询"
// @Param limit query int false "页数"
// @Param offset query int false "页码"
// @Param gender query string false "按性别查询"
// @Param tags query []string false "按标签查询"
// @Success 200 {object} model.BaseResponse{data=object{list=[]model.StudentInfo,total_count=int64}} "desc"
// @Router /v2/student [get]
func (h *stuHandle) QueryStudentsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.QueryStudentRequest
		c.ShouldBindQuery(&req)
		result, err := h.websvc.QueryStudent(req)

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
