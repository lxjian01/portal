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
	result := gorm.GetOrmDB().Table("user").Select("user_name","phone","email","weixin","update_user").Where("id = ?", m.Id).Updates(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteUser(id int) error {
	// 开始事务
	tx := gorm.GetOrmDB().Begin()
	defer tx.Rollback()
	// find user
	txUser := tx.Table("user").Where("id = ?", id)
	var user models.User
	txUser.First(&user)
	// delete user role
	result := tx.Table("user_role").Where("user_code = ?", user.UserCode).Delete(&models.Role{})
	if result.Error != nil {
		return result.Error
	}
	// delete user
	result = txUser.Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	// 提交事务
	return tx.Commit().Error
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