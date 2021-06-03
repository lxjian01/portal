package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorExporter(m *models.MonitorExporter) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_exporter").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorExporter(m *models.MonitorExporter) error {
	result := myorm.GetOrmDB().Table("monitor_exporter").Select("name","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorExporter(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_exporter").Where("id = ?", id).Delete(&models.MonitorExporter{})
	return result.RowsAffected, result.Error
}

func GetMonitorExporterList() (*[]models.MonitorExporterList, error) {
	dataList := make([]models.MonitorExporterList, 0)
	myorm.GetOrmDB().Table("monitor_exporter").Select("id","name").Find(&dataList)
	return &dataList, nil
}

func GetMonitorExporterPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorExporter, 0)
	tx := myorm.GetOrmDB().Table("monitor_exporter")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}