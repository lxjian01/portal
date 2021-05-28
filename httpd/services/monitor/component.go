package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorComponent(m *models.MonitorComponent) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_component").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorComponent(m *models.MonitorComponent) error {
	result := myorm.GetOrmDB().Table("monitor_component").Select("code","name","exporter","template","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorComponent(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_component").Where("id = ?", id).Delete(&models.MonitorComponent{})
	return result.RowsAffected, result.Error
}

func GetMonitorComponentList() (*[]models.MonitorComponentList, error) {
	dataList := make([]models.MonitorComponentList, 0)
	myorm.GetOrmDB().Table("monitor_component").Select("code","name").Find(&dataList)
	return &dataList, nil
}

func GetMonitorComponentPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorComponent, 0)
	tx := myorm.GetOrmDB().Table("monitor_component")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("code like ? or name like ? or exporter like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}