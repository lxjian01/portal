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

func UpdateRecordingRule(m *models.RecordingRuleAdd) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		var prometheusList []models.Prometheus
		err := tx.Table("prometheus").Where("id in ?", m.PrometheusIds).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// update recording rule
		err = tx.Table("recording_rule").Select("prometheus_id","name","expr").Where("id = ?", m.Id).Updates(m).Error
		if err != nil {
			return err
		}
		// find prometheus recording rule list
		oldPrometheusList := make([]prometheusRecordingRuleList, 0)
		err = myorm.GetOrmDB().Table("prometheus_recording_rule").Select("prometheus_recording_rule.recording_rule_id", "prometheus_recording_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_recording_rule.prometheus_id = prometheus.id").Where("prometheus_recording_rule.recording_rule_id = ?", m.Id).Find(&oldPrometheusList).Error
		if err != nil {
			return err
		}
		// delete prometheus recording rule
		err = tx.Table("prometheus_recording_rule").Where("recording_rule_id = ?", m.Id).Delete(&models.PrometheusRecordingRule{}).Error
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
		// delete consul Key/Value
		for _, item := range oldPrometheusList {
			code := item.PrometheusCode
			key := fmt.Sprintf("prometheus/%s/rules/recordings/%s", code, m.Record)
			client := consul.GetClient()
			_, consulErr := client.KV().Delete(key, nil)
			if consulErr != nil {
				log.Errorf("Delete prometheus %s rule recording error by %s", code, consulErr)
			}
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
	return err
}

func DeleteRecordingRule(id int) error {
	// delete recording rule
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		prometheusList := make([]prometheusRecordingRuleList, 0)
		err := myorm.GetOrmDB().Table("prometheus_recording_rule").Select("prometheus_recording_rule.recording_rule_id", "prometheus_recording_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_recording_rule.prometheus_id = prometheus.id").Where("prometheus_recording_rule.recording_rule_id = ?", id).Find(&prometheusList).Error
		if err != nil {
			return err
		}
		// delete prometheus recording rule
		err = tx.Table("prometheus_recording_rule").Where("recording_rule_id = ?", id).Delete(&models.PrometheusRecordingRule{}).Error
		if err != nil {
			return err
		}
		// delete recording rule
		recordingRule := models.RecordingRule{}
		tx.Table("recording_rule").Where("id = ?", id).Find(&recordingRule)
		err = tx.Table("recording_rule").Where("id = ?", id).Delete(&models.Prometheus{}).Error
		if err != nil {
			return err
		}

		// delete consul Key/Value
		for _, item := range prometheusList {
			code := item.PrometheusCode
			key := fmt.Sprintf("prometheus/%s/rules/recordings/%s", code, recordingRule.Record)
			client := consul.GetClient()
			_, consulErr := client.KV().Delete(key, nil)
			if consulErr != nil {
				log.Errorf("Delete prometheus %s rule recording error by %s", code, consulErr)
			}
		}
		return err
	})
	return err
}

type prometheusRecordingRuleList struct {
	RecordingRuleId int `gorm:"column:recording_rule_id" json:"recordingRuleId"`
	PrometheusId int `gorm:"column:prometheus_id" json:"prometheusId"`
	PrometheusCode string `gorm:"column:prometheus_code" json:"prometheusCode"`
	PrometheusName string `gorm:"column:prometheus_name" json:"prometheusName"`
}

func GetRecordingRulePage(pageIndex int, pageSize int, prometheusId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.RecordingRulePage, 0)
	tx := myorm.GetOrmDB().Table("recording_rule")
	tx.Select("recording_rule.id","recording_rule.name","recording_rule.record","recording_rule.expr","recording_rule.update_user","recording_rule.update_time")
	if prometheusId != 0 {
		tx.Joins("left join prometheus_recording_rule on prometheus_recording_rule.recording_rule_id = recording_rule.id")
		tx.Joins("left join prometheus on prometheus.id = prometheus_recording_rule.prometheus_id")
		tx.Where("prometheus.id = ?", prometheusId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("recording_rule.name like ? or recording_rule.record like ? or recording_rule.expr like ?", likeStr, likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	recordings := make([]prometheusRecordingRuleList, 0)
	myorm.GetOrmDB().Table("prometheus_recording_rule").Select("prometheus_recording_rule.recording_rule_id", "prometheus_recording_rule.prometheus_id","prometheus.code as prometheus_code","prometheus.name as prometheus_name").Joins("left join prometheus on prometheus_recording_rule.prometheus_id = prometheus.id").Find(&recordings)
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