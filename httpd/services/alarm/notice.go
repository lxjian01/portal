package alarm

import (
	"portal/global/myorm"
	"portal/httpd/models"
)

func AddAlarmSender(m *models.AlarmSender) (int, error) {
	err := myorm.GetOrmDB().Table("alarm_sender").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}