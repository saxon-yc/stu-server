package model

import (
	"time"

	"github.com/lib/pq"
)

// 租户表
type TenantDb struct {
	TenantID   uint32 `gorm:"column:tenant_id;autoIncrement;comment:主键" json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

type UserDb struct {
	// autoIncrement 不可与primaryKey、type同时使用否则不生效。
	// AUTO_INCREMENT 生效 gorm会自动根据字段类型设置数据库字段类型并设置为主键
	UserID     uint32    `gorm:"column:user_id;autoIncrement;comment:主键" json:"user_id"` // 写成AUTO_INCREMENT也可以
	TenantID   uint32    `gorm:"column:tenant_id;comment:租户ID" json:"tenant_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`

	UserName string         `gorm:"column:username" json:"username"`
	Password string         `json:"password"`
	Token    string         `json:"token"`
	Menus    pq.StringArray `gorm:"column:menus;type:text[];comment:菜单权限" json:"menus"` // 可用菜单
}

type StudentDb struct {
	ID         uint32         `gorm:"column:id;autoIncrement;comment:主键" json:"id"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime time.Time      `json:"update_time"`
	Name       string         `json:"name"`
	Age        int            `json:"age"`
	Gender     string         `json:"gender"` // male｜femal
	Address    string         `json:"address"`
	Tags       pq.StringArray `gorm:"column:tags;type:text[];comment:标签" json:"tags"`
}

type TagDb struct {
	ID         uint32    `gorm:"column:id;autoIncrement;comment:主键" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Label      string    `json:"label"`
	Content    string    `gorm:"column:content;type:text" json:"content"`
	Count      int       `json:"count"`
}
