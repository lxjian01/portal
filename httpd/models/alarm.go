package models

import "time"

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
	Record   string `gorm:"column:record;type:varchar(64)" json:"record" form:"record" binding:"required"`
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

type AlertingMetric struct {
	BaseModel
	Exporter   string `gorm:"column:exporter;type:varchar(64)" json:"exporter" form:"exporter" binding:"required"`
	Code   string `gorm:"column:code;type:varchar(64)" json:"code" form:"code" binding:"required"`
	Metric   string `gorm:"column:metric;type:varchar(256)" json:"metric" form:"metric" binding:"required"`
	Summary   string `gorm:"column:summary;type:varchar(128)" json:"summary" form:"summary" binding:"required"`
	Description   string `gorm:"column:description;type:varchar(256)" json:"description" form:"description" binding:"required"`
	Remark   string `gorm:"column:remark;type:varchar(512)" json:"remark" form:"remark" binding:""`
}

type AlertingMetricList struct {
	Id int `gorm:"column:id" json:"id"`
	Summary   string `gorm:"column:summary" json:"summary"`
}

type AlertingRule struct {
	BaseModel
	AlertingMetricId   int `gorm:"column:alerting_metric_id;type:int" json:"alertingMetricId" form:"alertingMetricId" binding:"required"`
	Operator   string `gorm:"column:operator;type:varchar(16)" json:"operator" form:"operator" binding:"required"`
	ThresholdValue   int `gorm:"column:threshold_value;type:int" json:"thresholdValue" form:"thresholdValue" binding:"required"`
	AlertingFor   string `gorm:"column:alerting_for;type:varchar(16)" json:"alertingFor" form:"alertingFor" binding:"required"`
	Severity   string `gorm:"column:severity;type:varchar(16)" json:"severity" form:"severity" binding:"required"`
}

type AlertingRuleAdd struct {
	AlertingRule
	PrometheusIds  []int         `gorm:"-" json:"prometheusIds" form:"prometheusIds" binding:"required"`
}

type AlertingRulePage struct {
	AlertingRule
	Exporter   string `gorm:"column:exporter" json:"exporter"`
	Code   string `gorm:"column:code" json:"code" form:"code"`
	Metric   string `gorm:"column:metric" json:"metric"`
	Summary   string `gorm:"column:summary" json:"summary"`
	Description   string `gorm:"column:description" json:"description"`
	PrometheusList []interface{} `json:"prometheusList"`
}

type PrometheusAlertingRule struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:""`
	PrometheusId   int `gorm:"column:prometheus_id;type:int" json:"prometheusId" form:"prometheusId" binding:"required"`
	AlertingRuleId   int `gorm:"column:alerting_rule_id;type:int" json:"alertingRuleId" form:"alertingRuleId" binding:"required"`
}

type AlarmNotice struct {
	Id        int `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id"`
	PrometheusCode   string `gorm:"column:prometheus_code;type:varchar(64)" json:"prometheusCode"`
	Exporter   string `gorm:"column:exporter;type:varchar(64)" json:"exporter"`
	MonitorResourceCode   string `gorm:"column:monitor_resource_code;type:varchar(64)" json:"monitorResourceCode"`
	Fingerprint   string `gorm:"column:fingerprint;type:varchar(64)" json:"fingerprint"`
	AlertName   string `gorm:"column:alert_name;type:varchar(128)" json:"alertName"`
	Instance   string `gorm:"column:instance;type:varchar(64)" json:"instance"`
	Summary   string `gorm:"column:summary;type:varchar(128)" json:"summary"`
	Description   string `gorm:"column:description;type:varchar(256)" json:"description"`
	Status   string `gorm:"column:status;type:varchar(16)" json:"status"`
	Severity   string `gorm:"column:severity;type:varchar(16)" json:"severity"`
	StartAt MyTime `gorm:"column:start_at;type:datetime" json:"startAt"`
	EndAt MyTime `gorm:"column:end_at;type:datetime" json:"endAt"`
	AlarmNumber   int `gorm:"column:alarm_number;type:int" json:"alarmNumber"`
	CreateTime MyTime `gorm:"column:create_time;type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime MyTime `gorm:"column:update_time;type:datetime;autoUpdateTime" json:"updateTime"`
}

type AlarmNoticePage struct {
	AlarmNotice
	PrometheusName   string `gorm:"column:prometheus_name" json:"prometheusName"`
	MonitorResourceName   string `gorm:"column:monitor_resource_name" json:"monitorResourceName"`
}

type Alert struct {
	Status string `json:"status"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Fingerprint string `json:"fingerprint"`
	StartsAt time.Time `json:"startsAt"`
	EndsAt time.Time `json:"endsAt"`
}

type AlertManagerNotice struct {
	Version string `json:"version"`
	GroupKey string `json:"groupkey"`
	Status string `json:"status"`
	Receiver string `json:"receiver"`
	TruncatedAlerts int `json:"truncatedAlerts"`
	GroupLabels map[string]string `json:"groupLabels"`
	CommonLabels map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalUrl string `json:"externalUrl"`
	Alerts []Alert `json:"alerts"`
}