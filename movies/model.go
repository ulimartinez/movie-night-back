package movies

import "movie-back/users"

type movieModel struct {
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"column:title"`
	Director string `gorm:"column:director"`
}

type MovieSubmissionModel struct {
	ID         uint `gorm:"primary_key"`
	MovieModel movieModel
	UserModel  users.UserModel
	Votes      uint `gorm:"column:votes"`
	Viewed     bool `gorm:"column:viewed"`
}
