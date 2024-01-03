package tagapi

import (
	"student-server/internal/dbsvc"
	"student-server/internal/websvc"

	"github.com/gin-gonic/gin"
)

/*
这段代码：是一个Handler接口类型，并将其初始化为nil的handler指针。这行代码的作用是确保handler结构体类型实现了Handler接口类型的所有方法，
如果没有实现，则会在编译时出现错误。这种技巧在Go语言中被称为interface check
*/
var _ TagHandler = (*tagHandle)(nil)

type TagHandler interface {
	QueryTagsAPI() gin.HandlerFunc
	CreateTagAPI() gin.HandlerFunc
	ChangeTagAPI() gin.HandlerFunc
	DeleteTagAPI() gin.HandlerFunc
}

type tagHandle struct {
	dbsvc  dbsvc.Servicer
	websvc websvc.WebServer

	// logger       *zap.Logger
	// cache        redis.Repo
	// hashids hash.Hash
}

func New(dbsvc dbsvc.Servicer, websvc websvc.WebServer) TagHandler {
	return &tagHandle{
		dbsvc:  dbsvc,
		websvc: websvc,
		// hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}
