package alarm

import (
	"errors"
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddRecordingRule(m *models.RecordingRuleAdd) (int, error) {
	prometheusUrl := ""
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// add recording rule
		err := tx.Table("recording_rule").Create(m).Error
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
		//// registry consul service
		//tags := make([]string, 0)
		//tags = append(tags, prometheus.Code)
		//tags = append(tags, monitorResource.Exporter)
		//meta := make(map[string]string, 0)
		//meta["resource"] = monitorResource.Code
		//client := consul.GetClient()
		//ip := strings.Split(m.Url,"/")[2]
		//tempIp := strings.Split(ip,":")
		//address := tempIp[0]
		//port,err := strconv.Atoi(tempIp[1])
		//if err != nil {
		//	return err
		//}
		//check := consulapi.AgentServiceCheck{
		//	HTTP:     m.Url,
		//	Interval: m.Interval,
		//}
		//serviceId := strconv.Itoa(m.Id)
		//registration := consulapi.AgentServiceRegistration{
		//	ID: serviceId,
		//	Name: m.Name,
		//	Address: address,
		//	Port: port,
		//	Tags: tags,
		//	Meta:  meta,
		//	Check: &check,
		//}
		//err = client.Agent().ServiceRegister(&registration)
		return err
	})

	// reload prometheus
	err = utils.PrometheusReload(prometheusUrl)
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

func GetRecordingRulePage(pageIndex int, pageSize int, prometheusId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.Prometheus, 0)
	tx := myorm.GetOrmDB().Table("recording_rule")
	tx.Select("recording_rule.id","recording_rule.prometheus_id","recording_rule.name","recording_rule.record","recording_rule.expr","prometheus.update_user","prometheus.update_time","prometheus.code as prometheus_code","prometheus.name as prometheus_name","prometheus.url as prometheus_url")
	tx.Joins("left join prometheus on prometheus.id = recording_rule.prometheus_id")
	if prometheusId != 0 {
		tx.Where("recording_rule.prometheus_id = ?", prometheusId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("recording_rule.name like ? or recording_rule.record like ? or recording_rule.expr like ?", likeStr, likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}