package nights

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"net/http"
)

func NightsRegister(router *gin.RouterGroup) {
	router.POST("/new", CreateNight)
}

func CreateNight(c *gin.Context) {
	//create the night
	validator := NewMovieNightValidator()
	err := validator.Bind(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	}
}
