package nights

import (
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/movies"
	"movie-back/users"
	"time"
)

type nightModel struct {
	gorm.Model
	Date  time.Time `gorm:"column:night_date"`
	Host  users.UserModel
	Movie movies.MovieSubmissionModel
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&nightModel{})
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
