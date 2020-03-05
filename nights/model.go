package nights

import (
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/movies"
	"movie-back/users"
	"time"
)

type NightModel struct {
	gorm.Model
	Date         time.Time `gorm:"column:night_date"`
	Host         users.UserModel
	HostID       uint
	GroupID      uint
	Movie        movies.MovieSubmissionModel
	SubmissionID uint
	Location     string
	History      bool
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&NightModel{})
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func SetMovie(data interface{}) (NightModel, error) {
	var nightModel NightModel
	var movieModel movies.MovieSubmissionModel
	type Result struct {
		max     uint
		GroupID uint
	}
	var result Result
	db := common.GetDB()
	db.Where(data).First(&nightModel)
	db.Table("movie_submission_models").Select("MAX(votes), group_id").Where("group_id = ?", nightModel.GroupID).Group("group_id").Scan(&result)
	db.Where(movies.MovieSubmissionModel{GroupID: result.GroupID, Votes: result.max}).First(&movieModel)
	err := db.Model(nightModel).Update(NightModel{SubmissionID: movieModel.ID}).Error
	return nightModel, err
}

func SetHistory(data interface{}) error {
	db := common.GetDB()
	return db.Model(data).Update(NightModel{History: true}).Error
}

func ListNights(data interface{}) ([]NightModel, error) {
	db := common.GetDB()
	var nights []NightModel
	err := db.Where(data).Find(&nights).Error
	return nights, err
}
