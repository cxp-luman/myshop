package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         int32     `gorm:"primarykey"`
	CreateAt   time.Time `gorm:"column:add_time"`
	UpdateTime time.Time `gorm:"column:modify_time"`
	DeleteAt   gorm.DeletedAt
}

type User struct {
	BaseModel
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(16);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Birthday *time.Time
	Gender   string `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
	Role     int    `gorm:"column:role;default:1;type:varchar(6) comment '1表示普通用户, 2表示管理员'"`
}
