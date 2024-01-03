package main

import (
	"student-server/config"
	"student-server/internal/dbsvc"
	router "student-server/internal/router"
	"student-server/internal/websvc"

	"github.com/spf13/viper"
)

func init() {
	config.New("/pitrix/config/stu_apiserver.yaml", "/pitrix/config/data_base.yaml")
	// config.New("config/stu_apiserver.yaml", "config/data_base.yaml")
}

// @title 学生管理系统
// @version 1.0
// @description 学生管理系统API文档
// @termsOfService http://swagger.io/terms/

// @contact.name saxon
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9999
// @BasePath /api
func main() {
	basicService := dbsvc.NewDbService()
	webService := websvc.NewWebService(basicService)
	r := router.NewRouter(basicService, webService)
	r.Run(viper.GetString("router.addr"))
}
