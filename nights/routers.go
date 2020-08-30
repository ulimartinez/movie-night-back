package nights

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
	"net/http"
	"strconv"
)

func NightsRegister(router *gin.RouterGroup) {
	router.POST("/new/:id", CreateNight)
	router.OPTIONS("/new/:id", preflight)
	router.POST("/set/:grid/:sid", SetMovieToNight)
	router.OPTIONS("/set/:grid/:sid", preflight)
	router.POST("/history/:id", MarkHistory)
	router.OPTIONS("/history/:id", preflight)
	router.GET("/list/:grid", FindNights)
	router.OPTIONS("/list/:grid", preflight)
	router.GET("/history/:grid", FindHistory)
}

func CreateNight(c *gin.Context) {
	//create the night
	validator := NewMovieNightValidator()
	err := validator.Bind(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	}
	myUser := c.MustGet("my_user_model").(users.UserModel)
	grid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	validator.MovieNight.Host = myUser
	validator.MovieNight.GroupID = uint(grid)
	err = SaveOne(&validator.MovieNight)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := NightSerializer{c, validator.MovieNight}
	c.JSON(http.StatusCreated, gin.H{"movienight": serializer.Response()})
}

func SetMovieToNight(c *gin.Context) {
	grid, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
	}
	sid, err := strconv.ParseUint(c.Param("sid"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
	}
	var nightModel NightModel
	nightModel.ID = uint(sid)
	nightModel.GroupID = uint(grid)
	newNight, err := SetMovie(&nightModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := NightSerializer{c, newNight}
	c.JSON(http.StatusOK, gin.H{"movienight": serializer.Response()})
}

func MarkHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
		return
	}
	var nightModel NightModel
	nightModel.ID = uint(id)
	err = SetHistory(&nightModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := NightSerializer{c, nightModel}
	c.JSON(http.StatusOK, gin.H{"movienight": serializer.Response()})
}

func FindNights(c *gin.Context) {
	grid, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
		return
	}
	nightModel := NightModel{GroupID: uint(grid), History: false}
	nights, err := ListNights(nightModel)
	serializer := NightsSerializer{c, nights}
	c.JSON(http.StatusOK, serializer.Response())
}

func FindHistory(c *gin.Context) {
	grid, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
		return
	}
	nightModel := NightModel{GroupID: uint(grid), History: true}
	nights, err := ListNights(nightModel)
	serializer := NightsSerializer{c, nights}
	c.JSON(http.StatusOK, serializer.Response())
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
