package websvc

import (
	"strconv"
	"strings"
	"student-server/internal/model"
)

var GenderMap = map[string]string{
	"male":   "男",
	"female": "女",
}

type StudentManager interface {
	CreateStudent(params model.CreateStudentRequest) error
	ChangeStudent(params model.ChangeStudentRequest) error
	DeleteStudent(id uint32) error
	ExportStudent(params model.ExportStudentRequest) (dataKey []map[string]string, data []map[string]interface{}, err error)
	QueryStudent(params model.QueryStudentRequest) (result model.QueryStudentResponse, err error)
}

func (svc *WebService) CreateStudent(params model.CreateStudentRequest) error {
	err := svc.dbService.CreateStudent(params)
	if err == nil {
		for _, v := range params.Tags {
			tag, err := svc.dbService.FindOneTagByLabel(v)
			if err == nil {
				s := strconv.FormatUint(uint64(tag.ID), 10)
				svc.dbService.IncrCountWithStuBindTag("label_" + s)
			}
		}
	}
	return err
}

func (svc *WebService) ChangeStudent(params model.ChangeStudentRequest) error {
	err := svc.dbService.ChangeStudent(params)
	return err
}

func (svc *WebService) QueryStudent(params model.QueryStudentRequest) (result model.QueryStudentResponse, err error) {
	result, err = svc.dbService.FindStudents(params)
	return
}
func (svc *WebService) ExportStudent(params model.ExportStudentRequest) (dataKey []map[string]string, data []map[string]interface{}, err error) {
	var tables model.QueryStudentResponse
	tables, err = svc.dbService.FindStudents(params.Query)
	if err != nil {
		return nil, nil, err
	}
	dataKey = make([]map[string]string, 0)

	// 生成表头
	for _, v := range params.Columns {
		switch v {
		case "id":
			dataKey = append(dataKey, map[string]string{"key": "id", "title": "学号", "width": "20", "is_num": "0"})
		case "name":
			dataKey = append(dataKey, map[string]string{"key": "name", "title": "姓名", "width": "20", "is_num": "0"})
		case "age":
			dataKey = append(dataKey, map[string]string{"key": "age", "title": "年龄", "width": "20", "is_num": "0"})
		case "gender":
			dataKey = append(dataKey, map[string]string{"key": "gender", "title": "性别", "width": "20", "is_num": "0"})
		case "tags":
			dataKey = append(dataKey, map[string]string{"key": "tags", "title": "标签", "width": "60", "is_num": "0"})
		case "address":
			dataKey = append(dataKey, map[string]string{"key": "address", "title": "家庭住址", "width": "60", "is_num": "0"})
		}
	}

	// 查询并组装数据
	data = make([]map[string]interface{}, 0)
	for _, v := range tables.List {
		data = append(data, map[string]interface{}{
			"id":      v.ID,
			"name":    v.Name,
			"age":     v.Age,
			"gender":  GenderMap[v.Gender],
			"tags":    strings.Join(v.Tags, "、"),
			"address": v.Address,
		})
	}
	return
}
func (svc *WebService) DeleteStudent(id uint32) error {
	stu := svc.dbService.FindOneStudentByID(id)
	for _, v := range stu.Tags {
		tag, err := svc.dbService.FindOneTagByLabel(v)
		if err == nil {
			s := strconv.FormatUint(uint64(tag.ID), 10)
			svc.dbService.DecrCountWithStuBindTag("label_" + s)
		}
	}
	err := svc.dbService.DeleteStudent(id)
	return err
}
