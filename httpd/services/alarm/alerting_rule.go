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

func GetAlertingRulePage(pageIndex int, pageSize int, exporter string, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlertingRulePage, 0)
	tx := myorm.GetOrmDB().Table("alerting_rule")
	tx.Select("id","exporter","alert","expr","for","severity","summary","description","update_user","update_time")
	if exporter != "" {
		tx.Where("exporter = ?", exporter)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("alert like ? or expr like ? or severity like ? or summary like ?", likeStr, likeStr, likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	alertings := make([]prometheusAlertingRuleList, 0)
	myorm.GetOrmDB().Table("prometheus_alerting_rule").Select("prometheus_alerting_rule.recording_rule_id", "prometheus_alerting_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_alerting_rule.prometheus_id = prometheus.id").Find(&alertings)
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
