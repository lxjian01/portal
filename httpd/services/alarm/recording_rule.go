package alarm

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddRecordingRule(m *models.RecordingRule) (int, error) {
	err := myorm.GetOrmDB().Table("recording_rule").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateRecordingRule(m *models.RecordingRule) error {
	result := myorm.GetOrmDB().Table("recording_rule").Select("prometheus_id","name","record","expr","remark").Where("id = ?", m.Id).Updates(m)
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
	tx.Select("recording_rule.id","recording_rule.prometheus_id","recording_rule.name","recording_rule.record","recording_rule.expr","recording_rule.remark","prometheus.update_user","prometheus.update_time","prometheus.code as prometheus_code","prometheus.name as prometheus_name","prometheus.url as prometheus_url")
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