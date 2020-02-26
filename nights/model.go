package nights

import (
	"github.com/jinzhu/gorm"
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
