package sysmgr

import (
	"errors"
	"portal/global/gorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMenu(m *models.Menu) (int, error) {
	err := gorm.GetOrmDB().Table("menu").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMenu(m *models.Menu) error {
	result := gorm.GetOrmDB().Table("menu").Select("pid","title","path","icon","sort","update_user","update_time").Where(`"id" = ?`, m.Id).Updates(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func DeleteMenu(id int) (int64, error) {
	result := gorm.GetOrmDB().Table("menu").Where("id = ?", id).Delete(&models.Menu{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetMenuDetail(id int) (*models.Menu, error) {
	var menu models.Menu
	gorm.GetOrmDB().Table("menu").Where("id = ?", id).First(&menu)
	return &menu, nil
}

func GetMenuList() (*[]models.Menu, error) {
	dataList := make([]models.Menu, 0)
	gorm.GetOrmDB().Table("menu").Select("id","pid","title","path","icon","sort").Find(&dataList)
	return &dataList, nil
}

func GetParentMenuList() (*[]models.Menu, error) {
	dataList := make([]models.Menu, 0)
	gorm.GetOrmDB().Table("menu").Select("id","pid","title").Where("pid = 0").Find(&dataList)
	return &dataList, nil
}

func GetMenuPage(pageIndex int, pageSize int, title string) (*utils.PageData, error) {
	dataList := make([]models.Menu, 0)
	tx := gorm.GetOrmDB().Table("menu")
	if title != "" {
		likeStr := "%" + title + "%"
		tx.Where("title like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	pageData.Data = &dataList
	return pageData, nil
}