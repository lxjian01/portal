package models

type AlarmGroup struct {
	BaseModel
	AlarmGroupName   string `gorm:"column:alarm_group_name;type:varchar(64);not null" json:"alarmGroupName" form:"alarmGroupName" binding:"required"`
}

type AlarmUser struct {
	BaseModel
	Name   string `gorm:"column:user_name;type:varchar(64);not null" json:"userName" form:"userName" binding:"required"`
	Phone   string `gorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:"required"`
	Email   string `gorm:"column:email;type:varchar(64);not null" json:"email" form:"email" binding:"required,email"`
	Weixin   string `gorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
}

type AlarmGroupUser struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id" binding:""`
	AlarmGroupId   string `gorm:"column:alarm_group_id;type:varchar(64)" json:"alarmGroupId" form:"alarmGroupId" binding:"required"`
	UserCode   string `gorm:"column:user_code;type:varchar(64)" json:"userCode" form:"userCode" binding:"required"`

}