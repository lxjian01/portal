package models

type AlarmGroup struct {
	BaseModel
	GroupName   string `gorm:"column:group_name;type:varchar(64);not null" json:"groupName" form:"groupName" binding:"required"`
}

type AlarmGroupAdd struct {
	AlarmGroup
	Users  []int         `gorm:"-" json:"users"`
}

type AlarmGroupPage struct {
	AlarmGroup
	Users  []string         `gorm:"_" json:"users"`
}

type AlarmUser struct {
	BaseModel
	UserName   string `gorm:"column:user_name;type:varchar(64);not null" json:"userName" form:"userName" binding:"required"`
	Phone   string `gorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:""`
	Email   string `gorm:"column:email;type:varchar(64);not null" json:"email" form:"email" binding:"required,email"`
	Weixin   string `gorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
}

type AlarmUserList struct {
	Id   int `json:"id"`
	UserName  string         `json:"userName"`
}

type AlarmGroupUser struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	UserId   int `gorm:"column:user_id;type:int" json:"userId" form:"userId" binding:"required"`
	GroupId   int `gorm:"column:group_id;type:int" json:"groupId" form:"groupId" binding:"required"`

}