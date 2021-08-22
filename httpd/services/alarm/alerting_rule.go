package alarm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"portal/global/consul"
	"portal/global/log"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddAlertingRule(m *models.AlertingRuleAdd) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		var prometheusList []models.Prometheus
		err := tx.Table("prometheus").Where("id in ?", m.PrometheusIds).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// find alerting metric
		var alertingMetric models.AlertingMetric
		err = tx.Table("alerting_metric").Where("id = ?", m.AlertingMetricId).Find(&alertingMetric).Error
		if err != nil {
			return err
		}
		// add alerting rule
		err = tx.Table("alerting_rule").Create(m).Error
		if err != nil {
			return err
		}
		// add prometheus alerting rule
		if len(m.PrometheusIds) <= 0 {
			return errors.New("Alerting rule必须要关联prometheus实例")
		}
		var prr []models.PrometheusAlertingRule
		for _, prometheusId := range m.PrometheusIds{
			var prometheusAlertingRule models.PrometheusAlertingRule
			prometheusAlertingRule.AlertingRuleId = m.Id
			prometheusAlertingRule.PrometheusId = prometheusId
			prr = append(prr, prometheusAlertingRule)
		}
		err = tx.Table("prometheus_alerting_rule").Create(&prr).Error
		if err != nil {
			return err
		}
		// registry consul Key/Value
		rule := utils.AlertingRule{}
		rule.Alert = fmt.Sprintf("%s_%d", alertingMetric.Code, m.Id)
		rule.Expr = fmt.Sprintf("%s%s%d", alertingMetric.Metric, m.Operator, m.ThresholdValue)
		rule.For = m.AlertingFor
		rule.Summary = alertingMetric.Summary
		rule.Description = alertingMetric.Description
		template, err := utils.GetAlertingRuleTemplate(&rule)
		if err != nil {
			return err
		}
		for _, item := range prometheusList {
			prometheusCode := item.Code
			key := fmt.Sprintf("prometheus/%s/rule/alertings/%s", prometheusCode, rule.Alert)
			consulErr := utils.PutConsul(key,template)
			if consulErr != nil {
				log.Errorf("Put prometheus %s alerting rule error by %s", prometheusCode, consulErr)
			}
		}
		return err
	})
	return m.Id, err
}

func UpdateAlertingRule(m *models.AlertingRuleAdd) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		var prometheusList []models.Prometheus
		err := tx.Table("prometheus").Where("id in ?", m.PrometheusIds).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// find alerting metric
		var alertingMetric models.AlertingMetric
		err = tx.Table("alerting_metric").Where("id = ?", m.AlertingMetricId).Find(&alertingMetric).Error
		if err != nil {
			return err
		}
		// update alerting rule
		err = tx.Table("alerting_rule").Select("alerting_metric_id","operator","threshold_value","alerting_for","severity").Where("id = ?", m.Id).Updates(m).Error
		if err != nil {
			return err
		}
		// find prometheus alerting rule list
		oldPrometheusList := make([]prometheusAlertingRuleList, 0)
		err = myorm.GetOrmDB().Table("prometheus_alerting_rule").Select("prometheus_alerting_rule.alerting_rule_id", "prometheus_alerting_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_alerting_rule.prometheus_id = prometheus.id").Where("prometheus_alerting_rule.alerting_rule_id = ?", m.Id).Find(&oldPrometheusList).Error
		if err != nil {
			return err
		}
		// delete prometheus alerting rule
		err = tx.Table("prometheus_alerting_rule").Where("alerting_rule_id = ?", m.Id).Delete(&models.PrometheusAlertingRule{}).Error
		if err != nil {
			return err
		}
		// add prometheus alerting rule
		if len(m.PrometheusIds) <= 0 {
			return errors.New("Alerting rule必须要关联prometheus实例")
		}
		var prr []models.PrometheusAlertingRule
		for _, prometheusId := range m.PrometheusIds{
			var prometheusAlertingRule models.PrometheusAlertingRule
			prometheusAlertingRule.AlertingRuleId = m.Id
			prometheusAlertingRule.PrometheusId = prometheusId
			prr = append(prr, prometheusAlertingRule)
		}
		err = tx.Table("prometheus_alerting_rule").Create(&prr).Error
		if err != nil {
			return err
		}

		rule := utils.AlertingRule{}
		rule.Alert = fmt.Sprintf("%s_%d", alertingMetric.Code, m.Id)
		rule.Expr = fmt.Sprintf("%s%s%d", alertingMetric.Metric, m.Operator, m.ThresholdValue)
		rule.For = m.AlertingFor
		rule.Summary = alertingMetric.Summary
		rule.Description = alertingMetric.Description
		// delete consul Key/Value
		for _, item := range oldPrometheusList {
			prometheusCode := item.PrometheusCode
			key := fmt.Sprintf("prometheus/%s/rule/alertings/%s", prometheusCode, rule.Alert)
			client := consul.GetClient()
			_, consulErr := client.KV().Delete(key, nil)
			if consulErr != nil {
				log.Errorf("Delete prometheus %s alerting rule error by %s", prometheusCode, consulErr)
			}
		}
		// registry consul Key/Value
		template, err := utils.GetAlertingRuleTemplate(&rule)
		if err != nil {
			return err
		}
		for _, item := range prometheusList {
			prometheusCode := item.Code
			key := fmt.Sprintf("prometheus/%s/rule/alertings/%s", prometheusCode, rule.Alert)
			consulErr := utils.PutConsul(key,template)
			if consulErr != nil {
				log.Errorf("Put prometheus %s alerting rule error by %s", prometheusCode, consulErr)
			}
		}
		return err
	})
	return err
}

func DeleteAlertingRule(id int) error {
	// delete alerting rule
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		prometheusList := make([]prometheusAlertingRuleList, 0)
		err := myorm.GetOrmDB().Table("prometheus_alerting_rule").Select("prometheus_alerting_rule.alerting_rule_id", "prometheus_alerting_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_alerting_rule.prometheus_id = prometheus.id").Where("prometheus_alerting_rule.alerting_rule_id = ?", id).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// delete prometheus alerting rule
		err = tx.Table("prometheus_alerting_rule").Where("alerting_rule_id = ?", id).Delete(&models.PrometheusAlertingRule{}).Error
		if err != nil {
			return err
		}
		// delete alerting rule
		alertingRule := models.AlertingRule{}
		tx.Table("alerting_rule").Where("id = ?", id).Find(&alertingRule)
		alertingMetric := models.AlertingMetric{}
		tx.Table("alerting_metric").Where("id = ?", alertingRule.AlertingMetricId).Find(&alertingMetric)
		err = tx.Table("alerting_rule").Where("id = ?", id).Delete(&models.AlertingRule{}).Error
		if err != nil {
			return err
		}

		alert := fmt.Sprintf("%s_%d", alertingMetric.Code, alertingRule.Id)
		// delete consul Key/Value
		for _, item := range prometheusList {
			prometheusCode := item.PrometheusCode
			key := fmt.Sprintf("prometheus/%s/rule/alertings/%s", prometheusCode, alert)
			client := consul.GetClient()
			_, consulErr := client.KV().Delete(key, nil)
			if consulErr != nil {
				log.Errorf("Delete prometheus %s alerting rule error by %s", prometheusCode, consulErr)
			}
		}
		return err
	})
	return err
}

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

