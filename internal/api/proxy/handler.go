package apiproxy

import (
	"student-server/internal/dbsvc"
	"student-server/internal/websvc"

	"github.com/gin-gonic/gin"
)

var _ ProxyHandler = (*proxyHandle)(nil)

type ProxyHandler interface {
	QuerySkus() gin.HandlerFunc
}

type proxyHandle struct {
	dbsvc  dbsvc.Servicer
	websvc websvc.WebServer
}

func New(dbsvc dbsvc.Servicer, websvc websvc.WebServer) ProxyHandler {
	return &proxyHandle{
		dbsvc:  dbsvc,
		websvc: websvc,
	}
}
