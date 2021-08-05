package alarm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"portal/global/consul"
	"portal/global/log"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func put(key string, value string) error {
	client := consul.GetClient()
	kv := client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return err
	}
	return nil
}

func AddRecordingRule(m *models.RecordingRuleAdd) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		var prometheusList []models.Prometheus
		err := tx.Table("prometheus").Where("id in ?", m.PrometheusIds).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// add recording rule
		err = tx.Table("recording_rule").Create(m).Error
		if err != nil {
			return err
		}
		// add prometheus recording rule
		if len(m.PrometheusIds) <= 0 {
			return errors.New("Recording rule必须要关联prometheus实例")
		}
		var prr []models.PrometheusRecordingRule
		for _, prometheusId := range m.PrometheusIds{
			var prometheusRecordingRule models.PrometheusRecordingRule
			prometheusRecordingRule.RecordingRuleId = m.Id
			prometheusRecordingRule.PrometheusId = prometheusId
			prr = append(prr, prometheusRecordingRule)
		}
		err = tx.Table("prometheus_recording_rule").Create(&prr).Error
		if err != nil {
			return err
		}
		// registry consul Key/Value
		rule := utils.RecordingRule{}
		rule.Record = m.Record
		rule.Expr = m.Expr
		template, err := utils.GetRecordingRuleTemplate(&rule)
		if err != nil {
			return err
		}
		for _, item := range prometheusList {
			code := item.Code
			key := fmt.Sprintf("prometheus/%s/rules/recordings/%s", code, m.Record)
			consulErr := put(key,template)
			if consulErr != nil {
				log.Errorf("Put prometheus %s rule recording error by %s", code, consulErr)
			}
		}
		return err
	})
	return m.Id, err
}

func UpdateRecordingRule(m *models.RecordingRule) error {
	result := myorm.GetOrmDB().Table("recording_rule").Select("prometheus_id","name","record","expr").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteRecordingRule(id int) (int64, error) {
	// delete recording rule
	result := myorm.GetOrmDB().Table("recording_rule").Where("id = ?", id).Delete(&models.Prometheus{})
	return result.RowsAffected, result.Error
}

type prometheusRecordingRuleList struct {
	RecordingRuleId int `gorm:"column:recording_rule_id" json:"recordingRuleId"`
	PrometheusId int `gorm:"column:prometheus_id" json:"prometheusId"`
	PrometheusName string `gorm:"column:prometheus_name" json:"prometheusName"`
}

func GetRecordingRulePage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.RecordingRulePage, 0)
	tx := myorm.GetOrmDB().Table("recording_rule")
	tx.Select("recording_rule.id","recording_rule.name","recording_rule.record","recording_rule.expr","recording_rule.update_user","recording_rule.update_time")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("recording_rule.name like ? or recording_rule.record like ? or recording_rule.expr like ?", likeStr, likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	recordings := make([]prometheusRecordingRuleList, 0)
	myorm.GetOrmDB().Table("prometheus_recording_rule").Select("prometheus_recording_rule.recording_rule_id", "prometheus_recording_rule.prometheus_id","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_recording_rule.prometheus_id = prometheus.id").Find(&recordings)
	for index, item := range dataList {
		for _, recording := range recordings {
			if item.Id == recording.RecordingRuleId {
				value := recording
				dataList[index].PrometheusList = append(dataList[index].PrometheusList, &value)
			}
		}
	}

	pageData.Data = &dataList
	return pageData, nil
}