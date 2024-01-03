package model

import "time"

type BaseResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ----------------------
// ---- admin manage ----
// ----------------------
type LoginRequest struct {
	UserName        string `form:"username,required"`
	Password        string `form:"password,required"`
	IsCreateAccount bool   `json:"is_registe"` // 用于标记在没有账户时，登录并且注册账户
}
type UserResponse struct {
	Token    string   `json:"token"`
	UserName string   `json:"username"`
	Menus    []string `json:"menus"` // 可用菜单
}
type DupuserRequest struct {
	UserName string `form:"username,required"`
}

// ------------------------
// ---- student manage ----
// ------------------------
type CreateStudentRequest struct {
	Name    string   `form:"name,required"`
	Age     int      `form:"age,required"`
	Gender  string   `form:"gender,required"` // male｜female
	Address string   `form:"address,required"`
	Tags    []string `form:"tags,omitempty"`
}
type ChangeStudentRequest struct {
	ID      uint32   `form:"id,required"`
	Age     int      `form:"age"`
	Address string   `form:"address"`
	Tags    []string `form:"tags,omitempty"`
}
type QueryStudentRequest struct {
	SearchWord string `form:"search_word"`
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
	Gender     string `form:"gender"`
	Tags       string `form:"tags"`
}

type ExportStudentRequest struct {
	Query   QueryStudentRequest `form:"query"`
	Columns []string            `form:"columns"`
}

type StudentInfo struct {
	ID         uint32    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	Tags       []string  `json:"tags"`
}
type QueryStudentResponse struct {
	TotalCount int64       `json:"total_count"`
	List       []StudentDb `json:"list"`
}
type DeleteStudentRequest struct {
	ID uint32 `uri:"id"`
}

// --------------------
// ---- tag manage ----
// --------------------
type CreateTagRequest struct {
	Label   string `form:"label"`
	Content string `form:"content"`
}
type ChangeTagRequest struct {
	ID      uint32 `form:"id"`
	Content string `form:"content"`
}
type DeleteTagRequest struct {
	ID uint32 `uri:"id"`
}
type QueryTagRequest struct {
	SearchWord string `form:"search_word"`
	StartTime  string `form:"start_time"`
	EndTime    string `form:"end_time"`
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
}
type QueryTagResponse struct {
	TotalCount int64   `json:"total_count"`
	List       []TagDb `json:"list"`
}
