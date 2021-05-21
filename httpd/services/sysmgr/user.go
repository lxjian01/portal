package sysmgr

import (
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddUser(m *models.User) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// add user role
		var userRoleList []models.UserRole
		for _, item := range m.Roles{
			var userRole models.UserRole
			userCode := m.UserCode
			roleCode := item
			userRole.UserCode = userCode
			userRole.RoleCode = roleCode
			userRoleList = append(userRoleList, userRole)
		}
		if len(userRoleList) > 0 {
			err := tx.Table("user_role").Create(&userRoleList).Error
			if err != nil {
				return err
			}
		}
		err := tx.Table("user").Create(m).Error
		return err
	})
	return m.Id, err
}

func UpdateUser(m *models.User) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// delete user role
		result := tx.Table("user_role").Where("user_code = ?", m.UserCode).Delete(&models.Role{})
		if result.Error != nil {
			return result.Error
		}
		// add user role
		var userRoleList []models.UserRole
		for _, item := range m.Roles{
			var userRole models.UserRole
			userCode := m.UserCode
			roleCode := item
			userRole.UserCode = userCode
			userRole.RoleCode = roleCode
			userRoleList = append(userRoleList, userRole)
		}
		if len(userRoleList) > 0 {
			err := tx.Table("user_role").Create(&userRoleList).Error
			if err != nil {
				return err
			}
		}
		// update user
		result = tx.Table("user").Select("user_name","phone","email","weixin","update_user").Where("id = ?", m.Id).Updates(m)
		return result.Error
	})
	return err
}

func DeleteUser(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
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
		return result.Error
	})
	return err
}

func GetUserDetail(id int) (*models.User, error) {
	var m models.User
	myorm.GetOrmDB().Table("user").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetUserList() (*[]models.User, error) {
	dataList := make([]models.User, 0)
	myorm.GetOrmDB().Table("user").Select("id","user_name","phone","email","weixin").Find(&dataList)
	return &dataList, nil
}

func GetUserPage(pageIndex int, pageSize int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.User, 0)
	tx := myorm.GetOrmDB().Table("user")
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