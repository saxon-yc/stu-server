package stuapi

import (
	"student-server/internal/dbsvc"
	"student-server/internal/websvc"

	"github.com/gin-gonic/gin"
)

var _ StuHandler = (*stuHandle)(nil)

type StuHandler interface {
	CreateStudentAPI() gin.HandlerFunc
	ChangeStudentAPI() gin.HandlerFunc
	QueryStudentsAPI() gin.HandlerFunc
	ExportStudentsAPI() gin.HandlerFunc
	ImportStudentsAPI() gin.HandlerFunc
	DeleteStudentAPI() gin.HandlerFunc
}

type stuHandle struct {
	dbsvc  dbsvc.Servicer
	websvc websvc.WebServer
}

func New(dbsvc dbsvc.Servicer, websvc websvc.WebServer) StuHandler {
	return &stuHandle{
		dbsvc:  dbsvc,
		websvc: websvc,
	}
}
