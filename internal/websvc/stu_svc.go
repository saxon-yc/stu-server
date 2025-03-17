package websvc

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"student-server/internal/model"
	errorcode "student-server/pkg/errors"
	"student-server/pkg/excelizelib"
	"student-server/pkg/util"
	"time"
)

var GenderMap = map[string]string{
	"male":   "男",
	"female": "女",
}

type StudentManager interface {
	CreateStudent(params model.CreateStudentRequest) error
	ChangeStudent(params model.ChangeStudentRequest) error
	DeleteStudent(id uint32) error
	ImportStudentsAPI(file *multipart.FileHeader) (total int64, err error)
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

/*
两种方案的对比:
方案	创建临时文件	内存直接处理
实现复杂度	简单（通用文件处理模式）	中等（需处理流式读取）
资源消耗	磁盘 I/O + 内存	仅内存（更高效）
适用场景	大文件处理（超过 100MB）	中小文件（< 100MB）
安全性	需要清理临时文件	无文件残留风险
性能瓶颈	磁盘写入速度	内存容量限制
结论：对于小于 100MB 的文件，推荐直接在内存中处理；对于超大文件仍需临时文件。
*/
func (svc *WebService) ImportStudentsAPI(fileHeader *multipart.FileHeader) (int64, error) {
	var total int64 = 0
	fmt.Printf("fileHeader: %v\n", fileHeader)
	if !util.IsExecl(fileHeader.Filename) {
		return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", errorcode.INVALID_FILE_UNCORRENT_MSG)
	}

	file, err := fileHeader.Open()
	fmt.Printf("file: %v\n", file)
	if err != nil {
		return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "文件读取失败")
	}
	defer file.Close()

	rows, err := excelizelib.ReadExcel(file)
	fmt.Printf("rows: %v\n", rows)
	if err != nil {
		return total, err
	}

	rowNum := 0
	var students []model.StudentDb

	for rows.Next() {
		rowNum++
		if rowNum == 1 {
			continue // 跳过标题行
		}

		row, err := rows.Columns()
		if err != nil {
			return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", fmt.Sprintf("第 %d 行读取失败", rowNum))
		}

		// // 数据校验与转换（示例）
		// if len(row) < 5 {
		// 	return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", fmt.Sprintf("第 %d 行数据不完整", rowNum))
		// }
		student := model.StudentDb{
			Name:       row[1],
			Gender:     row[3],
			Address:    row[4],
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		// 转换年龄字段
		if age, err := strconv.Atoi(row[2]); err == nil {
			student.Age = age
		} else {
			return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", fmt.Sprintf("第 %d 行年龄格式错误", rowNum))
		}
		// 转换性别字段
		if row[3] == "男" {
			student.Gender = "male"
		} else {
			student.Gender = "female"
		}

		students = append(students, student)
	}
	if len(students) == 0 {
		return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "无有效数据")
	}
	fmt.Printf("students: %v\n", students)
	total, err = svc.dbService.BatchInsertStudents(students)
	if err != nil {
		return total, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "批量插入失败")
	}

	return total, nil
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
