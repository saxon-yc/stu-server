package dbsvc

import (
	"fmt"
	"strconv"
	"strings"
	"student-server/internal/model"
	"time"

	"gorm.io/gorm"
)

func (b basicService) CreateStudent(params model.CreateStudentRequest) (err error) {
	err = b.gdb.Create(&model.StudentDb{
		Name:       params.Name,
		Age:        params.Age,
		Gender:     params.Gender,
		Address:    params.Address,
		Tags:       params.Tags,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}).Error
	return
}

func (b basicService) FindStudents(params model.QueryStudentRequest) (result model.QueryStudentResponse, err error) {
	var tx = b.gdb
	keyword := params.SearchWord
	if keyword != "" {
		// 字符串转换
		i, _ := strconv.Atoi(keyword)
		u := uint32(i)
		tx = tx.Where("name LIKE ? OR id= ? ", "%"+keyword+"%", u)
	}
	if params.Gender != "" {
		tx = tx.Where("gender= ?", params.Gender)
	}
	if params.Tags != "" {
		var tagSlice []string
		for _, v := range strings.Split(params.Tags, ",") {
			tagSlice = append(tagSlice, "'%"+v+"%'")
		}
		//Raw sql: SELECT * FROM student_db WHERE tags::text LIKE ALL (ARRAY['%花生过敏%', '%乳糖不耐受%']);
		sql := fmt.Sprintf("tags::text LIKE ALL (ARRAY[%s])", strings.Join(tagSlice, ", "))
		tx = tx.Where(sql)
	}
	result, err = findStudents(tx, params)
	return
}

func (b basicService) FindOneStudentByID(id uint32) (result model.StudentDb) {
	b.gdb.Where("id = ?", id).First(&result)
	return
}

func (b basicService) ChangeStudent(params model.ChangeStudentRequest) (err error) {
	err = b.gdb.Model(&model.StudentDb{ID: params.ID}).Updates(&model.StudentDb{UpdateTime: time.Now(), Age: params.Age, Tags: params.Tags, Address: params.Address}).Error
	return
}
func (b basicService) DeleteStudent(id uint32) (err error) {
	err = b.gdb.Delete(&model.StudentDb{}, id).Error
	return
}

func findStudents(tx *gorm.DB, params model.QueryStudentRequest) (result model.QueryStudentResponse, err error) {
	err = tx.
		Limit(params.Limit).
		Offset(params.Offset).
		Order("id asc").
		Find(&result.List).
		Count(&result.TotalCount).Error
	return
}
