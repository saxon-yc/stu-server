package router

import (
	adminapi "student-server/internal/api/admin"
	apiproxy "student-server/internal/api/proxy"
	stuapi "student-server/internal/api/student"
	tagapi "student-server/internal/api/tag"
	"student-server/internal/dbsvc"
	websvc "student-server/internal/websvc"
	"student-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouterAPI(rg *gin.RouterGroup, basicService dbsvc.Servicer, handlerService websvc.WebServer) {

	rv1 := rg.Group("/v1")
	adminAPI := adminapi.New(basicService, handlerService)
	{
		rv1.POST("/login", adminAPI.Login())
		rv1.GET("/dupuser", adminAPI.QueryUserExist())
	}

	// 需要 JWT 鉴权的路由：
	rv2 := rg.Group("/v2")
	// 使用JWT中间件
	rv2.Use(middleware.JWTAuthMiddleware())

	{
		rv2.GET("/user", adminAPI.QueryUserInfo())
	}

	stuAPI := stuapi.New(basicService, handlerService)
	{
		rv2.GET("/student", stuAPI.QueryStudentsAPI())
		rv2.POST("/student", stuAPI.CreateStudentAPI())
		rv2.PUT("/student", stuAPI.ChangeStudentAPI())
		rv2.POST("/student/export", stuAPI.ExportStudentsAPI())
		rv2.DELETE("/student/:id", stuAPI.DeleteStudentAPI())
	}

	tagAPI := tagapi.New(basicService, handlerService)
	{
		rv2.GET("/tag", tagAPI.QueryTagsAPI())
		rv2.POST("/tag", tagAPI.CreateTagAPI())
		rv2.PATCH("/tag", tagAPI.ChangeTagAPI())
		rv2.DELETE("/tag/:id", tagAPI.DeleteTagAPI())
	}

	proxyAPI := apiproxy.New(basicService, handlerService)
	{
		rv1.POST("/skus", proxyAPI.QuerySkus())
	}
}
