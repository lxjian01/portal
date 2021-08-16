package monitor

import (
	"errors"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddExporter(m *models.Exporter) (int, error) {
	err := myorm.GetOrmDB().Table("exporter").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateExporter(m *models.Exporter) error {
	result := myorm.GetOrmDB().Table("exporter").Select("code","exporter","git_url","remark").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteExporter(id int) (int64, error) {
	// find exporter
	var exporter models.Exporter
	err := myorm.GetOrmDB().Table("exporter").Where("id = ?", id).Find(&exporter).Error
	if err != nil {
		return 0, err
	}
	// monitor resource exists
	var monitorResourceCount int64
	err = myorm.GetOrmDB().Table("monitor_resource").Where("exporter = ?", exporter.Exporter).Count(&monitorResourceCount).Error
	if err != nil {
		return 0, err
	}
	if monitorResourceCount > 0 {
		return 0, errors.New("exporter下存在监控资源，不允许删除")
	}
	// alerting rule exists
	var alertingRuleCount int64
	err = myorm.GetOrmDB().Table("alerting_rule").Where("exporter = ?", exporter.Exporter).Count(&alertingRuleCount).Error
	if err != nil {
		return 0, err
	}
	if alertingRuleCount > 0 {
		return 0, errors.New("集群下存在alerting rule，不允许删除")
	}
	// delete exporter
	result := myorm.GetOrmDB().Table("exporter").Where("id = ?", id).Delete(&models.Exporter{})
	return result.RowsAffected, result.Error
}

func GetExporterList() (*[]models.ExporterList, error) {
	dataList := make([]models.ExporterList, 0)
	err := myorm.GetOrmDB().Table("exporter").Select("exporter").Find(&dataList).Error
	return &dataList, err
}

func GetExporterPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.Exporter, 0)
	tx := myorm.GetOrmDB().Table("exporter")
	tx.Select("id","exporter","git_url","remark","update_user","update_time")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("exporter like ? or git_url like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}

