package groups

import (
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/movies"
	"movie-back/users"
)

type GroupModel struct {
	ID          uint `gorm:"primary_key"`
	Friends     []users.UserModel
	Name        string
	Submissions []movies.MovieSubmissionModel `gorm:"foreignkey:GroupID"`
}
type GroupUserModel struct {
	gorm.Model
	UserModel    users.UserModel
	UserModelID  uint
	GroupModel   GroupModel
	GroupModelID uint
}

type GroupSubmissionsModel struct {
	gorm.Model
	GroupModel   GroupModel
	GroupModelID uint
	Submissions  []movies.MovieSubmissionModel
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

func GetGroups(userModel users.UserModel) ([]GroupModel, error) {
	var groups []GroupModel
	var userGroupModels []GroupUserModel
	db := common.GetDB()
	tx := db.Begin()
	tx.Where(&GroupUserModel{UserModelID: userModel.ID}).Find(&userGroupModels)

	for _, groupUserModel := range userGroupModels {
		var group GroupModel
		tx.Where(&GroupModel{ID: groupUserModel.GroupModelID}).First(&group)
		groups = append(groups, group)
	}
	err := tx.Commit().Error
	return groups, err
}

func GetUsers(groupModel GroupModel) ([]users.UserModel, error) {
	var usersArr []users.UserModel
	var groupUsers []GroupUserModel
	db := common.GetDB()
	tx := db.Begin()
	tx.Where(&GroupUserModel{GroupModelID: groupModel.ID}).Find(&groupUsers)
	for _, groupUser := range groupUsers {
		var user users.UserModel
		tx.Where(&users.UserModel{ID: groupUser.UserModelID}).First(&user)
		usersArr = append(usersArr, user)
	}
	err := tx.Commit().Error
	return usersArr, err
}
