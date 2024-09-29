package userdao

import (
	"fmt"
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func NewUserDao(db *gorm.DB) *Dao {
	_ = db.AutoMigrate(&User{})
	return &Dao{
		DB: db,
	}
}
func (dao *Dao) FindById(id string) (user *User, err error) {
	var usr User
	tx := dao.DB.First(&usr, id)
	if err = tx.Error; err != nil {
		return nil, fmt.Errorf("UserDao.FindById: %w", err)
	}
	return &usr, nil
}
