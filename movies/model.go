package movies

import (
	"movie-back/common"
	"movie-back/groups"
	"movie-back/users"
)

type MovieModel struct {
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"column:title"`
	Director string `gorm:"column:director"`
}

type MovieSubmissionModel struct {
	ID         uint `gorm:"primary_key"`
	MovieModel MovieModel
	UserModel  users.UserModel
	GroupModel groups.GroupModel
	UserID     uint
	GroupID    uint
	Votes      uint `gorm:"column:votes;DEFAULT:0"`
	Viewed     bool `gorm:"column:viewed;DEFAULT:0"`
}

func Automigrate() {
	db := common.GetDB()
	db.AutoMigrate(&MovieModel{})
	db.AutoMigrate(&MovieSubmissionModel{})
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func SubmitMovie(movie MovieModel, user users.UserModel, group groups.GroupModel) error {
	db := common.GetDB()
	submission := MovieSubmissionModel{
		ID:         0,
		MovieModel: movie,
		UserModel:  user,
		GroupModel: group,
		UserID:     user.ID,
		GroupID:    group.ID,
		Votes:      0,
		Viewed:     false,
	}
	err := db.Create(submission).Error
	return err
}
