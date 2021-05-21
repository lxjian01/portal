package sysmgr

import (
	"errors"
	"portal/global/gorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddUser(m *models.User) (int, error) {
	err := gorm.GetOrmDB().Table("user").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateUser(m *models.User) error {
	result := gorm.GetOrmDB().Table("user").Select("id","user_code","user_name","phone","email","weixin","update_user","update_time").Where("id = ?", m.Id).Updates(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteUser(id int) (int64, error) {
	result := gorm.GetOrmDB().Table("user").Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetUserDetail(id int) (*models.User, error) {
	var m models.User
	gorm.GetOrmDB().Table("user").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetUserList() (*[]models.User, error) {
	dataList := make([]models.User, 0)
	gorm.GetOrmDB().Table("user").Select("id","user_name","phone","email","weixin").Find(&dataList)
	return &dataList, nil
}

func GetUserPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.User, 0)
	tx := gorm.GetOrmDB().Table("user")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("user_code like ? or user_name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}