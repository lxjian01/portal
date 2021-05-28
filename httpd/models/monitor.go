package models

type MonitorCluster struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);not null" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64);not null" json:"name" form:"name" binding:"required"`
	PrometheusUrl   string `gorm:"column:prometheus_url;type:varchar(64);not null" json:"prometheusUrl" form:"prometheusUrl" binding:"required"`
}
