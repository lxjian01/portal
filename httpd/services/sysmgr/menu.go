package sysmgr

import (
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMenu(m *models.Menu) (int, error) {
	err := myorm.GetOrmDB().Table("menu").Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func UpdateMenu(m *models.Menu) error {
	result := myorm.GetOrmDB().Table("menu").Select("pid","title","path","icon","sort","update_user").Where("id = ?", m.Id).Updates(m)
	return result.Error
}

func DeleteMenu(id int) (int64, error) {
	result := myorm.GetOrmDB().Table("menu").Where("id = ?", id).Delete(&models.Menu{})
	return result.RowsAffected, result.Error
}

func GetMenuDetail(id int) (*models.Menu, error) {
	var m models.Menu
	myorm.GetOrmDB().Table("menu").Where("id = ?", id).First(&m)
	return &m, nil
}

func GetMenuList() (*[]models.Menu, error) {
	dataList := make([]models.Menu, 0)
	myorm.GetOrmDB().Table("menu").Select("id","pid","title","path","icon","sort").Order("pid desc,update_time asc,sort asc").Find(&dataList)
	return &dataList, nil
}

func GetParentMenuList() (*[]models.Menu, error) {
	dataList := make([]models.Menu, 0)
	myorm.GetOrmDB().Table("menu").Select("id","pid","title").Where("pid = 0").Find(&dataList)
	return &dataList, nil
}

func GetMenuPage(pageIndex int, pageSize int, title string) (*utils.PageData, error) {
	dataList := make([]models.Menu, 0)
	tx := myorm.GetOrmDB().Table("menu")
	if title != "" {
		likeStr := "%" + title + "%"
		tx.Where("title like ?", likeStr)
	}
	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	parentList := make([]models.Menu, 0)
	myorm.GetOrmDB().Table("menu").Select("id","pid","title").Where("pid = 0").Find(&parentList)
	for index, item := range dataList {
		for _, pitem := range parentList {
			if item.Pid == pitem.Id {
				value := pitem.Title
				dataList[index].PTitle = value
			}
		}
	}
	pageData.Data = &dataList
	return pageData, nil
}