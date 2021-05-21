package models

type Menu struct {
	BaseModel
	Pid         int `gorm:"column:pid;type:int;default=0" json:"pid" form:"pid"`
	Title   string `gorm:"column:title;type:varchar(64)" json:"title" form:"title" binding:"required"`
	PTitle   string `json:"ptitle"`
	Path      string `gorm:"column:path;type:varchar(64)" json:"path" form:"path" binding:"required"`
	Icon    string `gorm:"column:icon;type:varchar(64)" json:"icon" form:"icon" binding:""`
	Sort uint8 `gorm:"column:sort;type:tinyint;default=1" json:"sort" form:"sort" binding:""`
}

type User struct {
	BaseModel
	UserCode   string `gorm:"column:user_code;type:varchar(64);not null;uniqueIndex" json:"userCode" form:"userCode" binding:"required"`
	UserName   string `gorm:"column:user_name;type:varchar(64);not null" json:"userName" form:"userName" binding:"required"`
	Phone   string `gorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:"required"`
	Email   string `gorm:"column:email;type:varchar(64);not null" json:"email" form:"email" binding:"required,email"`
	Weixin   string `gorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
}

type Role struct {
	BaseModel
	RoleCode   string `gorm:"column:role_code;type:varchar(64)" json:"roleCode" form:"roleCode" binding:"required"`
	RoleName   string `gorm:"column:role_name;type:varchar(64)" json:"roleName" form:"roleName" binding:"required"`
}

type UserRole struct {
	BaseModel
	UserCode   string `gorm:"column:user_code;type:varchar(64)" json:"userCode" form:"userCode" binding:"required"`
	RoleCode   string `gorm:"column:role_code;type:varchar(64)" json:"roleCode" form:"roleCode" binding:"required"`
}
