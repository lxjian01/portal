package models

type AlarmGroup struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
}

type AlarmGroupAdd struct {
	AlarmGroup
	Users  []int         `gorm:"-" json:"users" form:"users" binding:""`
}

type AlarmGroupList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type AlarmGroupPage struct {
	AlarmGroup
	Users  []string         `gorm:"_" json:"users"`
}

type AlarmUser struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Phone   string `gorm:"column:phone;type:varchar(32)" json:"phone" form:"phone" binding:""`
	Email   string `gorm:"column:email;type:varchar(64)" json:"email" form:"email" binding:"required,email"`
	Weixin   string `gorm:"column:weixin;type:varchar(64)" json:"weixin" form:"weixin" binding:""`
}

type AlarmUserList struct {
	Id   int `json:"id"`
	Name  string         `json:"name"`
}

type AlarmGroupUser struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	AlarmUserId   int `gorm:"column:alarm_user_id;type:int" json:"alarmUserId" form:"alarmUserId" binding:"required"`
	AlarmGroupId   int `gorm:"column:alarm_group_id;type:int" json:"alarmGroupId" form:"alarmGroupId" binding:"required"`
}

type RecordingRule struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Record   string `gorm:"column:record;type:varchar(64);uniqueIndex" json:"record" form:"record" binding:"required"`
	Expr   string `gorm:"column:expr;type:varchar(64)" json:"expr" form:"expr" binding:"required"`
}

type RecordingRuleAdd struct {
	RecordingRule
	PrometheusIds  []int         `gorm:"-" json:"prometheusIds" form:"prometheusIds" binding:"required"`
}

type RecordingRulePage struct {
	RecordingRule
	PrometheusList []interface{} `json:"prometheusList"`
}

type PrometheusRecordingRule struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	PrometheusId   int `gorm:"column:prometheus_id;type:int" json:"prometheusId" form:"prometheusId" binding:"required"`
	RecordingRuleId   int `gorm:"column:recording_rule_id;type:int" json:"recordingRuleId" form:"recordingRuleId" binding:"required"`
}