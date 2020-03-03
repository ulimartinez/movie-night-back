package movies

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
	"net/http"
	"strconv"
)

func MovieRegister(router *gin.RouterGroup) {
	router.POST("submit/:grid", SubmitMovie)
	router.POST("vote/:id", VoteID)
	router.GET("list/:id", List)
}

func SubmitMovie(c *gin.Context) {
	myGroupID, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	movieValidator := NewMovieValidator()
	if err := movieValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := SaveOne(&movieValidator.MovieModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	sub, err := SubmitNewMovie(movieValidator.MovieModel, myUserModel, uint(myGroupID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := SubmissionSerializer{Submission: sub, c: c}
	c.JSON(http.StatusCreated, gin.H{"submission": serializer.Response()})
}

func VoteID(c *gin.Context) {
	movieID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
	}

	submissionm := MovieSubmissionModel{ID: uint(movieID)}
	submission, err := Vote(&submissionm)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := SubmissionSerializer{c, submission}
	c.JSON(http.StatusOK, gin.H{"submission": serializer.Response()})
}

func List(c *gin.Context) {
	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
	}
	submissions, err := ListGroupSubmissions(uint(groupID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := SubmissionsSerializer{c, submissions}
	c.JSON(http.StatusOK, gin.H{"submissions": serializer.Response()})
}
