package nights

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"time"
)

type MovieNightValidator struct {
	Night struct {
		Date time.Time `form:"datetime" json:"datetime" binding:"required" time_format:"2020-01-01 24:24"`
	} `json:"night"`
	MovieNight NightModel `json:"-"`
}

func (selfr *MovieNightValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, selfr)
	if err != nil {
		return err
	}
	selfr.MovieNight.Date = selfr.Night.Date
	return nil
}

func NewMovieNightValidator() MovieNightValidator {
	validator := MovieNightValidator{}
	return validator
}
