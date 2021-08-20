package alarm

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

type prometheusAlertingRuleList struct {
	AlertingRuleId int `gorm:"column:alerting_rule_id" json:"alertingRuleId"`
	PrometheusId int `gorm:"column:prometheus_id" json:"prometheusId"`
	PrometheusCode string `gorm:"column:prometheus_code" json:"prometheusCode"`
	PrometheusName string `gorm:"column:prometheus_name" json:"prometheusName"`
}

func GetAlertingRulePage(pageIndex int, pageSize int, prometheusId int, exporter string, alertingMetricId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlertingRulePage, 0)
	tx := myorm.GetOrmDB().Table("alerting_rule")
	tx.Select("alerting_rule.id","alerting_rule.alerting_metric_id","alerting_rule.operator","alerting_rule.threshold_value","alerting_rule.alerting_for","alerting_rule.severity","alerting_rule.update_user","alerting_rule.update_time","alerting_metric.exporter","alerting_metric.code","alerting_metric.metric","alerting_metric.summary","alerting_metric.description")
	if prometheusId != 0 {
		tx.Joins("left join prometheus_alerting_rule on prometheus_alerting_rule.alerting_rule_id = alerting_rule.id")
		tx.Joins("left join prometheus on prometheus.id = prometheus_alerting_rule.prometheus_id")
		tx.Where("prometheus.id = ?", prometheusId)
	}
	tx.Joins("left join alerting_metric on alerting_metric.id = alerting_rule.alerting_metric_id")
	if exporter != "" {
		tx.Where("alerting_metric.exporter = ?", exporter)
	}
	if alertingMetricId != 0 {
		tx.Where("alerting_rule.alerting_metric_id = ?", alertingMetricId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("alerting_metric.code like ? or alerting_metric.summary like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	alertings := make([]prometheusAlertingRuleList, 0)
	myorm.GetOrmDB().Table("prometheus_alerting_rule").Select("prometheus_alerting_rule.alerting_rule_id", "prometheus_alerting_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_alerting_rule.prometheus_id = prometheus.id").Find(&alertings)
	for index, item := range dataList {
		for _, alerting := range alertings {
			if item.Id == alerting.AlertingRuleId {
				value := alerting
				dataList[index].PrometheusList = append(dataList[index].PrometheusList, &value)
			}
		}
	}

	pageData.Data = &dataList
	return pageData, nil
}

