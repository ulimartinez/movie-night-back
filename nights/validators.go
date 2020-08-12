package nights

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"time"
)

type MovieNightValidator struct {
	Night struct {
		Date     time.Time `form:"date" json:"date" binding:"required" time_format:"2006-01-02"`
		Location string    `form:"location" json:"location" binding:"required"`
	} `json:"night"`
	MovieNight NightModel `json:"-"`
}

func (selfr *MovieNightValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, selfr)
	if err != nil {
		return err
	}
	selfr.MovieNight.Date = selfr.Night.Date
	selfr.MovieNight.Location = selfr.Night.Location
	return nil
}

func NewMovieNightValidator() MovieNightValidator {
	validator := MovieNightValidator{}
	return validator
}
