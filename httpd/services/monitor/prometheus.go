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
	result := myorm.GetOrmDB().Table("prometheus").Select("monitor_cluster_id","name","url","remark").Where("id = ?", m.Id).Updates(m)
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
	// delete prometheus
	result := myorm.GetOrmDB().Table("prometheus").Where("id = ?", id).Delete(&models.Prometheus{})
	return result.RowsAffected, result.Error
}

func GetPrometheusList(monitorClusterId int) (*[]models.PrometheusList, error) {
	dataList := make([]models.PrometheusList, 0)
	tx := myorm.GetOrmDB().Table("prometheus")
	tx.Select("id","name")
	if monitorClusterId != 0 {
		tx.Where("prometheus.monitor_cluster_id = ?", monitorClusterId)
	}
	tx.Find(&dataList)
	return &dataList, nil
}

func GetPrometheusPage(pageIndex int, pageSize int, monitorClusterId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.PrometheusPage, 0)
	tx := myorm.GetOrmDB().Table("prometheus")
	tx.Select("prometheus.id","prometheus.name","prometheus.url","prometheus.remark","prometheus.update_user","prometheus.update_time","prometheus.monitor_cluster_id","monitor_cluster.code as monitor_cluster_code","monitor_cluster.name as monitor_cluster_name")
	tx.Joins("left join monitor_cluster on monitor_cluster.id = prometheus.monitor_cluster_id")
	if monitorClusterId != 0 {
		tx.Where("prometheus.monitor_cluster_id = ?", monitorClusterId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("prometheus.name like ? or prometheus.url like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}

