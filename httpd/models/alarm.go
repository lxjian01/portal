package models

type AlarmGroup struct {
	BaseModel
	GroupName   string `gorm:"column:group_name;type:varchar(64);not null" json:"groupName" form:"groupName" binding:"required"`
}

type AlarmUser struct {
	BaseModel
	UserName   string `gorm:"column:user_name;type:varchar(64);not null" json:"userName" form:"userName" binding:"required"`
	Phone   string `gorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:"required"`
	Email   string `gorm:"column:email;type:varchar(64);not null" json:"email" form:"email" binding:"required,email"`
	Weixin   string `gorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
}

type AlarmGroupUser struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id" binding:""`
	UserId   string `gorm:"column:user_id;type:varchar(64)" json:"userId" form:"userId" binding:"required"`
	GroupId   string `gorm:"column:group_id;type:varchar(64)" json:"groupId" form:"groupId" binding:"required"`

}