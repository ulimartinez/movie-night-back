package nights

import (
	"fmt"
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
		Max     uint
		Groupid uint
	}
	var result Result
	db := common.GetDB()
	db.Where(data).First(&nightModel)
	db.Table("movie_submission_models").Select("MAX(votes) as max, group_id as groupid").Where("group_id = ?", nightModel.GroupID).Group("group_id").Scan(&result)
	db.Where(movies.MovieSubmissionModel{GroupID: result.Groupid, Votes: result.Max}).First(&movieModel)
	fmt.Print(result)
	nightModel.Movie = movieModel
	err := db.Model(nightModel).Update(NightModel{SubmissionID: movieModel.ID}).Error
	return nightModel, err
}

func SetHistory(data interface{}) error {
	db := common.GetDB()
	var night NightModel
	db.Where(data).Find(&night)
	db.Table("movie_submission_models").Where("id = ?", night.SubmissionID).Find(&night.Movie)
	db.Table("movie_submission_models").Where(night.Movie).Update(movies.MovieSubmissionModel{Viewed: true})
	return db.Model(data).Update(NightModel{History: true}).Error
}

func ListNights(data interface{}) ([]NightModel, error) {
	db := common.GetDB()
	var nights []NightModel
	realData := data.(NightModel)
	fmt.Print("%+v\n", data)
	err := db.Where(map[string]interface{}{"group_id": realData.GroupID, "History": realData.History}).Find(&nights).Error
	for i, night := range nights {
		db.Where("id = ?", night.SubmissionID).Find(&nights[i].Movie)
	}

	return nights, err
}
