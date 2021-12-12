package alarm

import (
	"portal/global/myorm"
	"portal/web/models"
	"portal/web/utils"
)

type AlertingMetricService struct {

}

func (service *AlertingMetricService) AddAlertingMetric(m *models.AlertingMetric) (int, error) {
	err := myorm.GetOrmDB().Table("alerting_metric").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func (service *AlertingMetricService) UpdateAlertingMetric(m *models.AlertingMetric) error {
	result := myorm.GetOrmDB().Table("alerting_metric").Select("exporter","code","metric","summary","description","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func (service *AlertingMetricService) DeleteAlertingMetric(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("alerting_metric").Where("id = ?", id).Delete(&models.AlertingMetric{})
	return result.RowsAffected, result.Error
}

func (service *AlertingMetricService) GetAlertingMetricDetail(id int) (*models.AlertingMetric, error) {
	var m models.AlertingMetric
	myorm.GetOrmDB().Table("alerting_metric").Where("id = ?", id).First(&m)
	return &m, nil
}

func (service *AlertingMetricService) GetAlertingMetricList(exporter string) (*[]models.AlertingMetricList, error) {
	dataList := make([]models.AlertingMetricList, 0)
	tx := myorm.GetOrmDB().Table("alerting_metric").Select("id","summary")
	if exporter != "" {
		tx.Where("exporter = ?", exporter)
	}
	tx.Find(&dataList)
	return &dataList, nil
}

func (service *AlertingMetricService) GetAlertingMetricPage(pageIndex int, pageSize int, exporter, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlertingMetric, 0)
	tx := myorm.GetOrmDB().Table("alerting_metric")
	tx.Select("id","exporter","code","metric","summary","description","remark","update_user","update_time")
	if exporter != "" {
		tx.Where("exporter = ?", exporter)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("code like ? or summary like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}