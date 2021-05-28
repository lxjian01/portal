package models

type MonitorCluster struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);not null" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64);not null" json:"name" form:"name" binding:"required"`
	PrometheusUrl   string `gorm:"column:prometheus_url;type:varchar(64);not null" json:"prometheusUrl" form:"prometheusUrl" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512);not null" json:"remark" form:"remark" binding:""`
}

type MonitorClusterList struct {
	BaseModel
	Code   string `json:"code"`
	Name   string `json:"name"`
}

type MonitorComponent struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);not null" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64);not null" json:"name" form:"name" binding:"required"`
	Template   string `gorm:"column:template;type:varchar(512);not null" json:"template" form:"template" binding:""`
	Remark   string `gorm:"column:remark;type:varchar(512);not null" json:"remark" form:"remark" binding:""`
}

type MonitorComponentList struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);not null" json:"code" form:"code" binding:"required"`
	Name   string `json:"name"`
}