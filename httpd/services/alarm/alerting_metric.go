package alarm

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddAlertingMetric(m *models.AlertingMetric) (int, error) {
	err := myorm.GetOrmDB().Table("alerting_metric").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateAlertingMetric(m *models.AlertingMetric) error {
	result := myorm.GetOrmDB().Table("alerting_metric").Select("exporter","code","name","metric","summary","description","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteAlertingMetric(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("alerting_metric").Where("id = ?", id).Delete(&models.AlertingMetric{})
	return result.RowsAffected, result.Error
}

func GetAlertingMetricList(exporter string) (*[]models.AlertingMetricList, error) {
	dataList := make([]models.AlertingMetricList, 0)
	tx := myorm.GetOrmDB().Table("alerting_metric").Select("id","name")
	if exporter != "" {
		tx.Where("exporter = ?", exporter)
	}
	tx.Find(&dataList)
	return &dataList, nil
}

func GetAlertingMetricPage(pageIndex int, pageSize int, exporter, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlertingMetric, 0)
	tx := myorm.GetOrmDB().Table("alerting_metric")
	tx.Select("id","exporter","code","name","metric","summary","description","remark","update_user","update_time")
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