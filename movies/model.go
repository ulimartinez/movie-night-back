package movies

import (
	"movie-back/common"
	"movie-back/users"
)

type movieModel struct {
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"column:title"`
	Director string `gorm:"column:director"`
}

type MovieSubmissionModel struct {
	ID         uint `gorm:"primary_key"`
	MovieModel movieModel
	UserModel  users.UserModel
	UserID     uint
	Votes      uint `gorm:"column:votes;DEFAULT:0"`
	Viewed     bool `gorm:"column:viewed"`
}

func Automigrate() {
	db := common.GetDB()
	db.AutoMigrate(&movieModel{})
	db.AutoMigrate(&MovieSubmissionModel{})
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
