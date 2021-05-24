package alarm

import (
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddAlarmGroup(m *models.AlarmGroup) (int, error) {
	err := myorm.GetOrmDB().Table("alarm_group").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateAlarmGroup(m *models.AlarmGroup) error {
	result := myorm.GetOrmDB().Table("alarm_group").Select("alarm_group_name").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteAlarmGroup(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find alarm group
		txAlarmGroup := tx.Table("alarm_group").Where("id = ?", id)
		var alarmGroup models.AlarmGroup
		txAlarmGroup.First(&alarmGroup)
		// delete alarm group user
		result := tx.Table("alarm_group_user").Where("alarm_group_id = ?", alarmGroup.Id).Delete(&models.AlarmGroup{})
		if result.Error != nil {
			return result.Error
		}
		// delete alarm group
		result = txAlarmGroup.Delete(&models.AlarmGroup{})
		return result.Error
	})
	return err


	result := myorm.GetOrmDB().Table("alarm_group").Where("id = ?", id).Delete(&models.AlarmGroup{})
	return result.Error
}

func GetAlarmGroupList() (*[]models.AlarmGroup, error) {
	dataList := make([]models.AlarmGroup, 0)
	myorm.GetOrmDB().Table("alarm_group").Select("id","alarm_group_name").Find(&dataList)
	return &dataList, nil
}

func GetAlarmGroupPage(pageIndex int, pageSize int, alarmGroupName string) (*utils.PageData, error) {
	dataList := make([]models.AlarmGroup, 0)
	tx := myorm.GetOrmDB().Table("alarm_group")
	if alarmGroupName != "" {
		likeStr := "%" + alarmGroupName + "%"
		tx.Where("alarm_group_name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}