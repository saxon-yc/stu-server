package dbsvc

import (
	"student-server/internal/model"
	"time"

	"gorm.io/gorm"
)

func (b basicService) CreateTag(params model.CreateTagRequest) (err error) {
	err = b.gdb.Create(&model.TagDb{
		Label:      params.Label,
		Content:    params.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}).Error
	return
}

func (b basicService) FindTags(params model.QueryTagRequest) (result model.QueryTagResponse, err error) {
	var tx *gorm.DB
	if params.SearchWord != "" {
		tx = b.gdb.Where("label LIKE ?", "%"+params.SearchWord+"%")
		result, err = findTags(tx, params)
		return
	}

	if params.StartTime != "" || params.EndTime != "" {
		tx = b.gdb.Where("create_time BETWEEN ? AND ?", params.StartTime, params.EndTime)
		result, err = findTags(tx, params)
		return
	}

	result, err = findTags(b.gdb, params)
	return
}
func (b basicService) FindOneTagByLabel(label string) (result model.TagDb, err error) {
	err = b.gdb.Where("label = ?", label).First(&result).Error
	return
}

func (b basicService) ChangeTag(params model.ChangeTagRequest) (err error) {
	err = b.gdb.Model(&model.TagDb{ID: params.ID}).
		Updates(map[string]interface{}{"update_time": time.Now(), "content": params.Content}).Error
	return
}
func (b basicService) DeleteTag(params model.DeleteTagRequest) (err error) {
	err = b.gdb.Delete(&model.TagDb{}, params.ID).Error
	return
}

func findTags(tx *gorm.DB, params model.QueryTagRequest) (result model.QueryTagResponse, err error) {
	err = tx.Limit(params.Limit).
		Offset(params.Offset).
		Order("create_time desc").
		Find(&result.List).
		Count(&result.TotalCount).Error
	return
}
