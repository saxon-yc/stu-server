package websvc

import (
	"strconv"
	"student-server/internal/model"
)

type TagManager interface {
	CreateTag(params model.CreateTagRequest) error
	ChangeTag(params model.ChangeTagRequest) error
	DeleteTag(params model.DeleteTagRequest) error
	QueryTag(params model.QueryTagRequest) (result model.QueryTagResponse, err error)
}

func (svc *WebService) CreateTag(params model.CreateTagRequest) error {
	err := svc.dbService.CreateTag(params)
	return err
}

func (svc *WebService) QueryTag(params model.QueryTagRequest) (result model.QueryTagResponse, err error) {
	tags, err := svc.dbService.FindTags(params)
	for _, v := range tags.List {
		s := strconv.FormatUint(uint64(v.ID), 10)
		str, _ := svc.dbService.FindCountByStuBindTag("label_" + s)

		i, _ := strconv.ParseInt(str, 10, 64)
		v.Count = int(i)
		result.List = append(result.List, v)
	}
	result.TotalCount = tags.TotalCount
	// for _, v := range result.List {
	// 	// svc.dbService.FindCountByStuBindTag(v.Label)
	// }
	return
}

func (svc *WebService) ChangeTag(params model.ChangeTagRequest) error {
	err := svc.dbService.ChangeTag(params)
	return err
}
func (svc *WebService) DeleteTag(params model.DeleteTagRequest) error {
	err := svc.dbService.DeleteTag(params)
	return err
}
