package sysmgr

import (
	"gorm.io/gorm"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddUser(m *models.User) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// add user
		err := tx.Table("user").Create(m).Error
		if err != nil {
			return err
		}
		// add user role
		if len(m.Roles) > 0 {
			var userRoleList []models.UserRole
			for _, item := range m.Roles{
				var userRole models.UserRole
				roleCode := item
				userRole.UserCode = m.UserCode
				userRole.RoleCode = roleCode
				userRoleList = append(userRoleList, userRole)
			}
			err = tx.Table("user_role").Create(&userRoleList).Error
		}
		return err
	})
	return m.Id, err
}

func UpdateUser(m *models.User) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// delete user role
		err := tx.Table("user_role").Where("user_code = ?", m.UserCode).Delete(&models.Role{}).Error
		if err != nil {
			return err
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
		err = tx.Table("user").Select("user_name","update_user").Where("id = ?", m.Id).Updates(m).Error
		return err
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
	var roles []role
	myorm.GetOrmDB().Table("user_role").Select("user_role.role_code","user_role.user_code", "role.role_name").Joins("left join role on user_role.role_code = role.role_code").Find(&roles)
	for _, pitem := range roles {
		if m.UserCode == pitem.UserCode {
			value := pitem
			m.Roles = append(m.Roles, value.RoleCode)
		}
	}
	return &m, nil
}

type role struct {
	UserCode  string
	RoleCode string
	RoleName string
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
	var roles []role
	myorm.GetOrmDB().Table("user_role").Select("user_role.role_code","user_role.user_code", "role.role_name").Joins("left join role on user_role.role_code = role.role_code").Find(&roles)
	for index, item := range dataList {
		for _, pitem := range roles {
			if item.UserCode == pitem.UserCode {
				value := pitem
				dataList[index].Roles = append(dataList[index].Roles, value.RoleName)
			}
		}
	}

	pageData.Data = &dataList
	return pageData, nil
}