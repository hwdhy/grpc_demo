package models

import (
	"gorm.io/gorm"
	"time"
)

// User 用户表
type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string // 用户名称
	Password  string //用户密码
	Salt      string // 密码盐值
	Type      string `gorm:"default:0"` // 用户类型 0:普通用户
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*User) TableName() string {
	return "user"
}
