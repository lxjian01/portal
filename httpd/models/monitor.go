package models

type MonitorCluster struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64);uniqueIndex" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	PrometheusUrl   string `gorm:"column:prometheus_url;type:varchar(64)" json:"prometheusUrl" form:"prometheusUrl" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorClusterList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type MonitorPrometheus struct {
	BaseModel
	MonitorClusterId   int `gorm:"column:monitor_cluster_id;type:int" json:"monitorClusterId" form:"monitorClusterId" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	PrometheusUrl   string `gorm:"column:prometheus_url;type:varchar(128);uniqueIndex" json:"prometheusUrl" form:"prometheusUrl" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorPrometheusList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type MonitorPrometheusPage struct {
	MonitorPrometheus
	MonitorClusterCode   string `gorm:"column:monitor_cluster_code" json:"monitorClusterCode"`
	MonitorClusterName   string `gorm:"column:monitor_cluster_name" json:"monitorClusterName"`
}

type MonitorComponent struct {
	BaseModel
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(64)" json:"name" form:"name" binding:"required"`
	Exporter   string `gorm:"column:exporter;type:varchar(64)" json:"exporter" form:"exporter" binding:"required"`
	GithubUrl   string `gorm:"column:github_url;type:varchar(64)" json:"githubUrl" form:"githubUrl" binding:"required"`
	Template   string `gorm:"column:template;type:varchar(512)" json:"template" form:"template" binding:""`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type MonitorComponentList struct {
	Id int `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
}

type MonitorTarget struct {
	BaseModel
	MonitorClusterId   int `gorm:"column:monitor_cluster_id;type:int" json:"monitorClusterId" form:"monitorClusterId" binding:"required"`
	MonitorComponentId   int `gorm:"column:monitor_component_id;type:int" json:"monitorComponentId" form:"monitorComponentId" binding:"required"`
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
	MonitorClusterCode   string `gorm:"column:monitor_cluster_code" json:"monitorClusterCode"`
	MonitorClusterName   string `gorm:"column:monitor_cluster_name" json:"monitorClusterName"`
	MonitorComponentCode   string `gorm:"column:monitor_component_code" json:"monitorComponentCode"`
	MonitorComponentName   string `gorm:"column:monitor_component_name" json:"monitorComponentName"`
	Exporter   string `gorm:"column:exporter" json:"exporter"`
	GroupList []interface{} `json:"alarmGroupList"`
}

type MonitorTargetAlarmGroup struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	MonitorTargetId   int `gorm:"column:monitor_target_id;type:int" json:"monitorTargetId" form:"monitorTargetId" binding:"required"`
	AlarmGroupId   int `gorm:"column:alarm_group_id;type:int" json:"alarmGroupId" form:"alarmGroupId" binding:"required"`
}
