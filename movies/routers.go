package movies

import (
	"errors"
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
	"net/http"
	"strconv"
	"time"
)

func MovieRegister(router *gin.RouterGroup) {
	router.POST("submit/:grid", SubmitMovie)
	router.POST("vote/:id", VoteID)
	router.GET("list/:id", List)
}

func SubmitMovie(c *gin.Context) {
	GroupID, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	myGroupID := uint(GroupID)
	if myGroupID == 0 {
		myGroupID = c.MustGet("my_group_id").(uint)
	}
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
	sub, err := SubmitNewMovie(movieValidator.MovieModel, myUserModel, myGroupID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := SubmissionSerializer{Submission: sub, C: c}
	c.JSON(http.StatusCreated, gin.H{"submission": serializer.Response()})
}

func VoteID(c *gin.Context) {
	myUser := c.MustGet("my_user_model").(users.UserModel)
	if myUser.LastVote.AddDate(0, 0, 1).After(time.Now()) {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("date", errors.New("no votes available yet")))
		return
	}
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
	err = users.UpdateVoted(myUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
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
