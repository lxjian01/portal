package sysmgr

import (
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
	return result.Error
}

func DeleteRole(id int) error {
	// 开始事务
	tx := gorm.GetOrmDB().Begin()
	// find role
	txRole := tx.Table("role").Where("id = ?", id)
	var role models.Role
	txRole.First(&role)
	// delete user role
	result := tx.Table("user_role").Where("role_code = ?", role.RoleCode).Delete(&models.Role{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// delete role
	result = txRole.Delete(&models.Role{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// 提交事务
	return tx.Commit().Error
}

func GetRoleDetail(id int) (*models.Role, error) {
	var m models.Role
	gorm.GetOrmDB().Table("role").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetRoleList() (*[]models.Role, error) {
	dataList := make([]models.Role, 0)
	gorm.GetOrmDB().Table("role").Select("id","role_code","role_name").Find(&dataList)
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