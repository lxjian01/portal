package sysmgr

import (
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddRole(m *models.Role) (int, error) {
	err := myorm.GetOrmDB().Table("role").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateRole(m *models.Role) error {
	result := myorm.GetOrmDB().Table("role").Select("role_name").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteRole(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find role
		txRole := tx.Table("role").Where("id = ?", id)
		var role models.Role
		txRole.First(&role)
		// delete role user
		result := tx.Table("user_role").Where("role_code = ?", role.RoleCode).Delete(&models.Role{})
		if result.Error != nil {
			return result.Error
		}
		// delete role
		result = txRole.Delete(&models.Role{})
		return result.Error
	})
	return err
}

func GetRoleList() (*[]models.Role, error) {
	dataList := make([]models.Role, 0)
	myorm.GetOrmDB().Table("role").Select("id","role_code","role_name").Find(&dataList)
	return &dataList, nil
}

func GetRolePage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.Role, 0)
	tx := myorm.GetOrmDB().Table("role")
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("role_code like ? or role_name like ?", likeStr, likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}