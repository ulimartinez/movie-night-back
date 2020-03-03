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
