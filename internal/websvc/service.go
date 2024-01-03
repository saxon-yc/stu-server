package websvc

import (
	"student-server/internal/dbsvc"
)

/*
这段代码：是一个Handler接口类型，并将其初始化为nil的handler指针。这行代码的作用是确保handler结构体类型实现了Handler接口类型的所有方法，
如果没有实现，则会在编译时出现错误。这种技巧在Go语言中被称为interface check
*/
var _ WebServer = (*WebService)(nil)

type WebServer interface {
	AdminManager
	TagManager
	StudentManager
}

type WebService struct {
	dbService dbsvc.Servicer
}

func NewWebService(dbService dbsvc.Servicer) WebServer {
	return &WebService{
		dbService: dbService,
	}
}
