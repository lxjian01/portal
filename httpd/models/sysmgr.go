package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Menu struct {
	BaseModel
	Pid         int `myorm:"column:pid;type:int;default=0" json:"pid" form:"pid"`
	Title   string `myorm:"column:title;type:varchar(64)" json:"title" form:"title" binding:"required"`
	PTitle   string `myorm:"-" json:"ptitle"`
	Path      string `myorm:"column:path;type:varchar(64)" json:"path" form:"path" binding:"required"`
	Icon    string `myorm:"column:icon;type:varchar(64)" json:"icon" form:"icon" binding:""`
	Sort uint8 `myorm:"column:sort;type:tinyint;default=1" json:"sort" form:"sort" binding:""`
}

type User struct {
	BaseModel
	UserCode   string `myorm:"column:user_code;type:varchar(64);not null;uniqueIndex" json:"userCode" form:"userCode" binding:"required"`
	UserName   string `myorm:"column:user_name;type:varchar(64);not null" json:"userName" form:"userName" binding:"required"`
	Phone   string `myorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:"required"`
	Email   string `myorm:"column:email;type:varchar(64);not null" json:"email" form:"email" binding:"required,email"`
	Weixin   string `myorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
	Roles  []string               `myorm:"-" json:"roles"`
}

type Roles []string

// Value 实现方法
func (m Roles) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan 实现方法
func (m *Roles) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &m)
}


type Role struct {
	BaseModel
	RoleCode   string `myorm:"column:role_code;type:varchar(64);not null;uniqueIndex" json:"roleCode" form:"roleCode" binding:"required"`
	RoleName   string `myorm:"column:role_name;type:varchar(64);not null" json:"roleName" form:"roleName" binding:"required"`
}

type UserRole struct {
	Id        int `myorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id" binding:""`
	UserCode   string `myorm:"column:user_code;type:varchar(64)" json:"userCode" form:"userCode" binding:"required"`
	RoleCode   string `myorm:"column:role_code;type:varchar(64)" json:"roleCode" form:"roleCode" binding:"required"`
}
