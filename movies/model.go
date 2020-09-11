package movies

import (
	"movie-back/common"
	"movie-back/users"
)

type MovieModel struct {
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"column:title"`
	Director string `gorm:"column:director"`
	Year	 uint	`gorm:"column:year"`
}

type MovieSubmissionModel struct {
	ID         uint `gorm:"primary_key"`
	MovieModel MovieModel
	UserModel  users.UserModel
	UserID     uint
	GroupID    uint
	MovieID    uint `gorm:"foreign_key:MovieID"`
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

func SubmitNewMovie(movie MovieModel, user users.UserModel, groupID uint) (MovieSubmissionModel, error) {
	db := common.GetDB()
	submission := MovieSubmissionModel{
		ID:         0,
		MovieModel: movie,
		UserModel:  user,
		UserID:     user.ID,
		GroupID:    groupID,
		MovieID:    movie.ID,
		Votes:      0,
		Viewed:     false,
	}
	err := db.Create(&submission).Error
	return submission, err
}

func Vote(condition interface{}) (MovieSubmissionModel, error) {
	var submission MovieSubmissionModel
	db := common.GetDB()
	db.Where(condition).First(&submission)
	err := db.Model(&submission).Update(MovieSubmissionModel{Votes: submission.Votes + 1}).Error
	return submission, err
}
func ListGroupSubmissions(condition uint) ([]MovieSubmissionModel, error) {
	db := common.GetDB()
	var submissions []MovieSubmissionModel
	err := db.Where("group_id = ? AND viewed = ?", condition, false).Find(&submissions).Error
	return submissions, err
}
func GetMovie(condition interface{}) (MovieModel, error) {
	db := common.GetDB()
	var movie MovieModel
	err := db.Where(condition).First(&movie).Error
	return movie, err
}
