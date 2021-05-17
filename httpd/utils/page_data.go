package utils

import (
	"errors"
	"gorm.io/gorm"
)

type PageData struct {
	Data interface{} `json:"results"`
	Total int64 `json:"count"`
	PageIndex int `json:"page_index"`
	PageSize int `json:"page_size"`
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