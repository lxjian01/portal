package sysmgr

import (
	"errors"
	"portal/global/gorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddRole(m *models.Role) (int, error) {
	err := gorm.GetOrmDB().Table("role").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateRole(m *models.Role) error {
	result := gorm.GetOrmDB().Table("role").Select("role_name").Where("id = ?", m.Id).Updates(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteRole(id int) (int64, error) {
	result := gorm.GetOrmDB().Table("role").Where("id = ?", id).Delete(&models.Role{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetRoleDetail(id int) (*models.Role, error) {
	var m models.Role
	gorm.GetOrmDB().Table("role").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetRoleList() (*[]models.Role, error) {
	dataList := make([]models.Role, 0)
	gorm.GetOrmDB().Table("role").Select("id","role_name").Find(&dataList)
	return &dataList, nil
}

func GetRolePage(pageIndex int, pageSize int, roleName string) (*utils.PageData, error) {
	dataList := make([]models.Role, 0)
	tx := gorm.GetOrmDB().Table("role")
	if roleName != "" {
		likeStr := "%" + roleName + "%"
		tx.Where("role_name like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}