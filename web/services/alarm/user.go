package alarm

import (
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/web/models"
	"portal/web/utils"
)

func AddAlarmUser(m *models.AlarmUser) (int, error) {
	err := myorm.GetOrmDB().Table("alarm_user").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateAlarmUser(m *models.AlarmUser) error {
	result := myorm.GetOrmDB().Table("alarm_user").Select("name","phone","email","weixin").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteAlarmUser(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find alarm group
		txAlarmUser := tx.Table("alarm_user").Where("id = ?", id)
		var alarmUser models.AlarmUser
		txAlarmUser.First(&alarmUser)
		// delete alarm group user
		result := tx.Table("alarm_group_user").Where("user_id = ?", alarmUser.Id).Delete(&models.AlarmUser{})
		if result.Error != nil {
			return result.Error
		}
		// delete alarm group
		result = txAlarmUser.Delete(&models.AlarmUser{})
		return result.Error
	})
	return err
}

func GetAlarmUserList() (*[]models.AlarmUserList, error) {
	dataList := make([]models.AlarmUserList, 0)
	myorm.GetOrmDB().Table("alarm_user").Select("id","name").Find(&dataList)
	return &dataList, nil
}

func GetAlarmUserPage(pageIndex int, pageSize int, name string) (*utils.PageData, error) {
	dataList := make([]models.AlarmUser, 0)
	tx := myorm.GetOrmDB().Table("alarm_user")
	if name != "" {
		likeStr := "%" + name + "%"
		tx.Where("name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}