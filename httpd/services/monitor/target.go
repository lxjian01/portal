package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorTarget(m *models.MonitorTarget) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_target").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorTarget(m *models.MonitorTarget) error {
	result := myorm.GetOrmDB().Table("monitor_target").Select("code","name","prometheus_url","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorTarget(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_target").Where("id = ?", id).Delete(&models.MonitorTarget{})
	return result.RowsAffected, result.Error
}

func GetMonitorTargetPage(pageIndex int, pageSize int, monitorClusterId int, monitorComponentId int, alarmGroupId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorTarget, 0)
	tx := myorm.GetOrmDB().Table("monitor_target")
	if monitorClusterId != 0 {
		tx.Where("monitor_cluster_id = ?", monitorClusterId)
	}
	if monitorComponentId != 0 {
		tx.Where("monitor_component_id = ?", monitorComponentId)
	}
	if alarmGroupId != 0 {
		tx.Where("alarm_group_id = ?", alarmGroupId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("name like ? or url like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}