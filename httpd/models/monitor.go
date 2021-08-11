package models

type Prometheus struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);uniqueIndex" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Url   string `gorm:"column:url;type:varchar(128);uniqueIndex" json:"url" form:"url" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type PrometheusList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type Exporter struct {
	BaseModel
	Exporter   string `gorm:"column:exporter;type:varchar(64);uniqueIndex" json:"exporter" form:"exporter" binding:"required"`
	GitUrl   string `gorm:"column:git_url;type:varchar(64);uniqueIndex" json:"gitUrl" form:"gitUrl" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type ExporterList struct {
	Exporter   string `gorm:"column:exporter" json:"exporter"`
}

type MonitorResource struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Exporter   string `gorm:"column:exporter;type:varchar(64)" json:"exporter" form:"exporter" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorResourceList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type MonitorTarget struct {
	BaseModel
	PrometheusId   int `gorm:"column:prometheus_id;type:int" json:"prometheusId" form:"prometheusId" binding:"required"`
	MonitorResourceId   int `gorm:"column:monitor_resource_id;type:int" json:"monitorResourceId" form:"monitorResourceId" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Url   string `gorm:"column:url;type:varchar(128)" json:"url" form:"url" binding:""`
	Interval   string `gorm:"column:interval;type:varchar(32)" json:"interval" form:"interval" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorTargetAdd struct {
	MonitorTarget
	AlarmGroupIds  []int         `gorm:"-" json:"alarmGroupIds" form:"alarmGroupIds" binding:"required"`
}

type MonitorTargetPage struct {
	MonitorTarget
	PrometheusCode   string `gorm:"column:prometheus_code" json:"prometheusCode"`
	PrometheusName   string `gorm:"column:prometheus_name" json:"prometheusName"`
	PrometheusUrl   string `gorm:"column:prometheus_url" json:"prometheusUrl"`
	MonitorResourceCode   string `gorm:"column:monitor_resource_code" json:"monitorResourceCode"`
	MonitorResourceName   string `gorm:"column:monitor_resource_name" json:"monitorResourceName"`
	Exporter   string `gorm:"column:exporter" json:"exporter"`
	GroupList []interface{} `json:"alarmGroupList"`
}

type MonitorTargetAlarmGroup struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	MonitorTargetId   int `gorm:"column:monitor_target_id;type:int" json:"monitorTargetId" form:"monitorTargetId" binding:"required"`
	AlarmGroupId   int `gorm:"column:alarm_group_id;type:int" json:"alarmGroupId" form:"alarmGroupId" binding:"required"`
}
