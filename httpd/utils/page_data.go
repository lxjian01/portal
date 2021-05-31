package utils

import (
	"errors"
	"gorm.io/gorm"
)

type PageData struct {
	Data interface{} `json:"data"`
	Total int64 `json:"total"`
	PageIndex int `json:"pageIndex"`
	PageSize int `json:"pageSize"`
}

func GetPageData(tx *gorm.DB, pageIndex int, pageSize int,data interface{}) (*PageData, error) {
	if pageIndex < 0 {
		return nil, errors.New("pageIndex mast greater 0")
	}
	if pageSize > 100 {
		return nil, errors.New("pageSize mast 1 - 100")
	}
	offset := (pageIndex - 1) * pageSize
	var total int64
	tx.Count(&total)
	tx.Limit(pageSize).Offset(offset).Find(data)
	pageData := &PageData{}
	pageData.Data = data
	pageData.Total = total
	return pageData, nil


}