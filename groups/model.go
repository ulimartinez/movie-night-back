package groups

import (
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/users"
)

type GroupModel struct {
	ID      uint `gorm:"primary_key"`
	Friends []users.UserModel
	Name    string
}
type GroupUserModel struct {
	gorm.Model
	UserModel    users.UserModel
	UserModelID  uint
	GroupModel   GroupModel
	GroupModelID uint
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&GroupModel{})
	db.AutoMigrate(&GroupUserModel{})
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneGroup(condition interface{}) (GroupModel, error) {
	db := common.GetDB()
	var groupModel GroupModel
	err := db.Where(condition).First(&groupModel).Error
	return groupModel, err
}
func (group GroupModel) AddToGroup(user users.UserModel) error {
	db := common.GetDB()
	err := db.Create(&GroupUserModel{
		UserModelID:  user.ID,
		GroupModelID: group.ID,
	}).Error
	return err
}
