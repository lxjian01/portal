package monitor

import (
	"errors"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddPrometheus(m *models.Prometheus) (int, error) {
	err := myorm.GetOrmDB().Table("prometheus").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdatePrometheus(m *models.Prometheus) error {
	result := myorm.GetOrmDB().Table("prometheus").Select("code","name","url","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeletePrometheus(id int) (int64, error) {
	// monitor target exists
	var monitorTargetCount int64
	err := myorm.GetOrmDB().Table("monitor_target").Where("prometheus_id = ?", id).Count(&monitorTargetCount).Error
	if err != nil {
		return 0, err
	}
	if monitorTargetCount > 0 {
		return 0, errors.New("集群下存在监控目标，不允许删除")
	}
	// recording rule exists
	var recordingRuleCount int64
	err = myorm.GetOrmDB().Table("recording_rule").Where("prometheus_id = ?", id).Count(&recordingRuleCount).Error
	if err != nil {
		return 0, err
	}
	if recordingRuleCount > 0 {
		return 0, errors.New("集群下存在recording rule，不允许删除")
	}
	// alerting rule exists
	var alertingRuleCount int64
	err = myorm.GetOrmDB().Table("alerting_rule").Where("prometheus_id = ?", id).Count(&alertingRuleCount).Error
	if err != nil {
		return 0, err
	}
	if alertingRuleCount > 0 {
		return 0, errors.New("集群下存在alerting rule，不允许删除")
	}
	// delete prometheus
	result := myorm.GetOrmDB().Table("prometheus").Where("id = ?", id).Delete(&models.Prometheus{})
	return result.RowsAffected, result.Error
}

func GetPrometheusList() (*[]models.PrometheusList, error) {
	dataList := make([]models.PrometheusList, 0)
	err := myorm.GetOrmDB().Table("prometheus").Select("id","code","name").Find(&dataList).Error
	return &dataList, err
}

func GetPrometheusPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.Prometheus, 0)
	tx := myorm.GetOrmDB().Table("prometheus")
	tx.Select("prometheus.id","prometheus.code","prometheus.name","prometheus.url","prometheus.remark","prometheus.update_user","prometheus.update_time")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("prometheus.code like ? or prometheus.name like ? or prometheus.url like ?", likeStr, likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}

