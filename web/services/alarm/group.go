package alarm

import (
	"errors"
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/web/models"
	"portal/web/utils"
)

func AddAlarmGroup(m *models.AlarmGroupAdd) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// add alarm group
		err := tx.Table("alarm_group").Create(m).Error
		if err != nil {
			return err
		}
		// add alarm group user
		if len(m.Users) > 0 {
			var alarmGroupUserList []models.AlarmGroupUser
			for _, item := range m.Users{
				var alarmGroupUser models.AlarmGroupUser
				userId := item
				alarmGroupUser.AlarmGroupId = m.Id
				alarmGroupUser.AlarmUserId = userId
				alarmGroupUserList = append(alarmGroupUserList, alarmGroupUser)
			}
			err = tx.Table("alarm_group_user").Create(&alarmGroupUserList).Error
		}
		return err
	})
	return m.Id, err
}

func UpdateAlarmGroup(m *models.AlarmGroupAdd) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// delete alarm group user
		err := tx.Table("alarm_group_user").Where("alarm_group_id = ?", m.Id).Delete(&models.AlarmGroupUser{}).Error
		if err != nil {
			return err
		}
		// add alarm group user
		if len(m.Users) > 0 {
			var alarmGroupUserList []models.AlarmGroupUser
			for _, item := range m.Users{
				var alarmGroupUser models.AlarmGroupUser
				userId := item
				alarmGroupUser.AlarmGroupId = m.Id
				alarmGroupUser.AlarmUserId = userId
				alarmGroupUserList = append(alarmGroupUserList, alarmGroupUser)
			}
			err = tx.Table("alarm_group_user").Create(&alarmGroupUserList).Error
		}
		// update alarm group
		err = tx.Table("alarm_group").Select("name").Where("id = ?", m.Id).Updates(m).Error
		return err
	})
	return err
}

func DeleteAlarmGroup(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// monitor target alert group exists
		var monitorTargetAlarmGroupCount int64
		err := myorm.GetOrmDB().Table("monitor_target_alarm_group").Where("alarm_group_id = ?", id).Count(&monitorTargetAlarmGroupCount).Error
		if err != nil {
			return err
		}
		if monitorTargetAlarmGroupCount > 0 {
			return errors.New("告警组下存在监控指标，不允许删除")
		}
		// delete alarm group user
		result := tx.Table("alarm_group_user").Where("alarm_group_id = ?", id).Delete(&models.AlarmGroup{})
		if result.Error != nil {
			return result.Error
		}
		// delete alarm group
		result = tx.Table("alarm_group").Where("id = ?", id).Delete(&models.AlarmGroup{})
		return result.Error
	})
	return err
}

type user struct {
	Id  int
	AlarmUserId int
	AlarmUserName string
	AlarmGroupId int
}

func GetAlarmGroupDetail(id int) (*models.AlarmGroupAdd, error) {
	var m models.AlarmGroupAdd
	myorm.GetOrmDB().Table("alarm_group").Where("id = ?", id).First(&m)
	var users []user
	myorm.GetOrmDB().Table("alarm_group_user").Select("alarm_group_user.id","alarm_group_user.alarm_group_id","alarm_group_user.alarm_user_id","alarm_user.name as alarm_user_name").Joins("left join alarm_user on alarm_group_user.alarm_user_id = alarm_user.id").Find(&users)
	for _, u := range users {
		value := u.AlarmGroupId
		if m.Id == value {
			m.Users = append(m.Users, u.AlarmUserId)
		}
	}
	return &m, nil
}

func GetAlarmGroupList() (*[]models.AlarmGroupList, error) {
	dataList := make([]models.AlarmGroupList, 0)
	myorm.GetOrmDB().Table("alarm_group").Select("id","name").Find(&dataList)
	return &dataList, nil
}

func GetAlarmGroupPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.AlarmGroupPage, 0)
	tx := myorm.GetOrmDB().Table("alarm_group")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	var users []user
	myorm.GetOrmDB().Table("alarm_group_user").Select("alarm_group_user.id","alarm_group_user.alarm_group_id","alarm_group_user.alarm_user_id","alarm_user.name as alarm_user_name").Joins("left join alarm_user on alarm_group_user.alarm_user_id = alarm_user.id").Find(&users)
	for index, item := range dataList {
		for _, pitem := range users {
			if item.Id == pitem.AlarmGroupId {
				value := pitem
				dataList[index].Users = append(dataList[index].Users, value.AlarmUserName)
			}
		}
	}

	pageData.Data = &dataList
	return pageData, nil
}