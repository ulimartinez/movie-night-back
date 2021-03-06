package movies

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
)

type MovieValidator struct {
	Movie struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Director string `form:"director" json:"director"`
		Year     uint   `form:"year" json:"year"`
	} `json:"movie"`
	MovieModel MovieModel `json:"-"`
}

func (self *MovieValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.MovieModel.Title = self.Movie.Name
	self.MovieModel.Year = self.Movie.Year
	return nil
}
func NewMovieValidator() MovieValidator {
	movieValidator := MovieValidator{}
	return movieValidator
}
