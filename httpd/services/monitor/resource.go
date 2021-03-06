package monitor

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorResource(m *models.MonitorResource) (int, error) {
	err := myorm.GetOrmDB().Table("monitor_resource").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMonitorResource(m *models.MonitorResource) error {
	result := myorm.GetOrmDB().Table("monitor_resource").Select("code","name","exporter","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMonitorResource(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("monitor_resource").Where("id = ?", id).Delete(&models.MonitorResource{})
	return result.RowsAffected, result.Error
}

func GetMonitorResourceList() (*[]models.MonitorResourceList, error) {
	dataList := make([]models.MonitorResourceList, 0)
	myorm.GetOrmDB().Table("monitor_resource").Select("id","code","name").Find(&dataList)
	return &dataList, nil
}

func GetMonitorResourcePage(pageIndex int, pageSize int, exporter, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorResource, 0)
	tx := myorm.GetOrmDB().Table("monitor_resource")
	tx.Select("id","code","name","exporter","remark","update_user","update_time")
	if exporter != "" {
		tx.Where("exporter = ?", exporter)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("code like ? or name like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}