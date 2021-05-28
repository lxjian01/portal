package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorCluster(m *models.MonitorCluster) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_cluster").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorCluster(m *models.MonitorCluster) error {
	result := myorm.GetOrmDB().Table("monitor_cluster").Select("name","prometheus_url").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorCluster(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_cluster").Where("id = ?", id).Delete(&models.MonitorCluster{})
	return result.RowsAffected, result.Error
}

func GetMonitorClusterList() (*[]models.MonitorCluster, error) {
	dataList := make([]models.MonitorCluster, 0)
	myorm.GetOrmDB().Table("monitor_cluster").Select("code","name").Find(&dataList)
	return &dataList, nil
}

func GetMonitorClusterPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorCluster, 0)
	tx := myorm.GetOrmDB().Table("monitor_cluster")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("code like ? or name like ? or prometheus_url like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}