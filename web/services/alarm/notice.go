package alarm

import (
	"portal/global/myorm"
	"portal/web/models"
	"portal/web/utils"
)

func AddOrUpdateAlarmNotice(m *models.AlarmNotice) error {
	var err error
	var alarmNotice models.AlarmNotice
	myorm.GetOrmDB().Table("alarm_notice").Where("alert_name = ? and instance = ? and severity = ? and status != 'resolved'", m.AlertName, m.Instance, m.Severity).Find(&alarmNotice)
	if alarmNotice.Id == 0 {
		err = myorm.GetOrmDB().Table("alarm_notice").Create(m).Error
	}else {
		m.AlarmNumber = alarmNotice.AlarmNumber + 1
		err = myorm.GetOrmDB().Table("alarm_notice").Select("summary","description","status","severity","end_at","alarm_number").Where("id = ?", alarmNotice.Id).Updates(&m).Error
	}
	// 异步发送告警通知
	return err
}

func GetAlarmNoticePage(pageIndex int, pageSize int, prometheusCode string, monitorResourceCode string, severity string, status string, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlarmNoticePage, 0)
	tx := myorm.GetOrmDB().Table("alarm_notice")
	tx.Select("alarm_notice.id","alarm_notice.prometheus_code","alarm_notice.alert_name","alarm_notice.instance","alarm_notice.summary","alarm_notice.description","alarm_notice.status","alarm_notice.severity","alarm_notice.start_at","alarm_notice.end_at","alarm_notice.alarm_number","alarm_notice.labels","alarm_notice.create_time","alarm_notice.update_time","prometheus.name as prometheus_name")
	tx.Joins("left join prometheus on prometheus.code = alarm_notice.prometheus_code")
	if prometheusCode != "" {
		tx.Where("alarm_notice.prometheus_code = ?", prometheusCode)
	}
	if severity != "" {
		tx.Where("alarm_notice.severity = ?", severity)
	}
	if status != "" {
		tx.Where("alarm_notice.status = ?", status)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("alarm_notice.alert_name like ? or alarm_notice.instance like ? or alarm_notice.summary like ?", likeStr, likeStr, likeStr)
	}

	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)

	return pageData, err
}

