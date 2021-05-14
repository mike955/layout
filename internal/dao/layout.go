package dao

import "gorm.io/gorm"

type Layout struct {
	CommonTimeModel
	Id        uint64 `json:"id" gorm"primaryKey"`
	layout    string `json:"layout"`
	IsDeleted uint64 `json:"is_deleted" gorm:"default:0" `
}

type LayoutDao struct {
	DB *gorm.DB
}

func NewLayoutDao() *LayoutDao {
	return &LayoutDao{
		DB: DB,
	}
}

func (dao *LayoutDao) Create(data Layout) (err error) {
	err = DB.Create(&data).Error
	return
}
