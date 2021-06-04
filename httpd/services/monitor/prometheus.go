package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorPrometheus(m *models.MonitorPrometheus) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_prometheus").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorPrometheus(m *models.MonitorPrometheus) error {
	result := myorm.GetOrmDB().Table("monitor_prometheus").Select("monitor_cluster_id","name","prometheus_url","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorPrometheus(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_prometheus").Where("id = ?", id).Delete(&models.MonitorPrometheus{})
	return result.RowsAffected, result.Error
}

func GetMonitorPrometheusList() (*[]models.MonitorPrometheusList, error) {
	dataList := make([]models.MonitorPrometheusList, 0)
	myorm.GetOrmDB().Table("monitor_prometheus").Select("id","name").Find(&dataList)
	return &dataList, nil
}

func GetMonitorPrometheusPage(pageIndex int, pageSize int, monitorClusterId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorPrometheusPage, 0)
	tx := myorm.GetOrmDB().Table("monitor_prometheus")
	tx.Select("monitor_prometheus.id","monitor_prometheus.name","monitor_prometheus.prometheus_url","monitor_prometheus.remark","monitor_prometheus.create_user","monitor_prometheus.create_time","monitor_prometheus.update_user","monitor_prometheus.update_time","monitor_prometheus.monitor_cluster_id","monitor_cluster.code as monitor_cluster_code","monitor_cluster.name as monitor_cluster_name")
	tx.Joins("left join monitor_cluster on monitor_prometheus.monitor_cluster_id = monitor_cluster.id")
	if monitorClusterId != 0 {
		tx.Where("monitor_cluster_id = ?", monitorClusterId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("name like ? or prometheus_url like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}