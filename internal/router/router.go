package router

import (
	"student-server/docs"
	"student-server/internal/dbsvc"
	websvc "student-server/internal/websvc"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(basicService dbsvc.Servicer, handlerService websvc.WebServer) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	pprof.Register(r)                  // pprof 性能优化
	corsConfig := cors.DefaultConfig() // cors：gin的中间件，主要处理跨域
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig)) // 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// swaggerTag := flag.Bool("swagger", true, "Whether to generate swagger document at build time")

	// flag.Parse()
	// 使用flag控制是否生成swagger，默认开启
	if viper.GetBool("swagger.enable") {
		// if swaggerTag != nil && *swaggerTag {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	rg := r.Group(viper.GetString("router.basePath"))
	basePath := rg.BasePath()
	docs.SwaggerInfo.BasePath = basePath

	SetRouterAPI(rg, basicService, handlerService)

	return r
}
