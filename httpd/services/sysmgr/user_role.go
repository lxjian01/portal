package sysmgr

import (
	"errors"
	"portal/global/gorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddUserRole(m *models.UserRole) (int, error) {
	err := gorm.GetOrmDB().Table("user_role").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateUserRole(m *models.UserRole) error {
	result := gorm.GetOrmDB().Table("user_role").Select("user_code", "role_code").Where("id = ?", m.Id).Updates(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteUserRole(id int) (int64, error) {
	result := gorm.GetOrmDB().Table("user_role").Where("id = ?", id).Delete(&models.Role{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetUserRoleDetail(id int) (*models.UserRole, error) {
	var m models.UserRole
	gorm.GetOrmDB().Table("user_role").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetUserRolePage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.UserRole, 0)
	tx := gorm.GetOrmDB().Table("user_role")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("user_code like ? or role_code like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}