package models

type MonitorCluster struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	PrometheusUrl   string `gorm:"column:prometheus_url;type:varchar(64)" json:"prometheusUrl" form:"prometheusUrl" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorClusterList struct {
	BaseModel
	Code   string `json:"code"`
	Name   string `json:"name"`
}

type MonitorComponent struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Exporter   string `gorm:"column:exporter;type:varchar(64)" json:"exporter" form:"exporter" binding:"required"`
	Template   string `gorm:"column:template;type:varchar(512)" json:"template" form:"template" binding:""`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorComponentList struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Name   string `json:"name"`
}

type MonitorTarget struct {
	BaseModel
	MonitorClusterId   int `gorm:"column:monitor_cluster_id;type:int" json:"monitorClusterId" form:"monitorClusterId" binding:"required"`
	MonitorComponentId   int `gorm:"column:monitor_component_id;type:int" json:"monitorComponentId" form:"monitorComponentId" binding:"required"`
	AlarmGroupId   int `gorm:"column:alarm_group_id;type:int" json:"alarmGroupId" form:"alarmGroupId" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Interval   int `gorm:"column:interval;type:int" json:"interval" form:"interval" binding:"required"`
	Url   string `gorm:"column:Url;type:varchar(128)" json:"Url" form:"Url" binding:""`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}