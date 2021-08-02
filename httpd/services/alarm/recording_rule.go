package alarm

import (
	"portal/global/myorm"
	"portal/httpd/models"
)

func AddRecordingRule(m *models.RecordingRule) (int, error) {
	err := myorm.GetOrmDB().Table("recording_rule").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateRecordingRule(m *models.RecordingRule) error {
	result := myorm.GetOrmDB().Table("recording_rule").Select("monitor_cluster_id").Where("id = ?", m.Id).Updates(m)
	return result.Error
}