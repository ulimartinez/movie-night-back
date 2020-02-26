package groups

import (
	"errors"
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
	"net/http"
	"strconv"
)

func GroupCreate(router *gin.RouterGroup) {
	router.POST("/create", GroupCreation)
	router.POST("/groupadd/:id", UserAddGroup)
}

func GroupCreation(c *gin.Context) {
	groupValidator := NewGroupValidator()
	if err := groupValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := SaveOne(&groupValidator.groupModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_group_model", groupValidator.groupModel)
	UserAddGroup(c)
}
func UserAddGroup(c *gin.Context) {
	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(groupID)
	if id == 0 {
		id = c.MustGet("my_group_model").(GroupModel).ID
	}
	groupModel, err := FindOneGroup(&GroupModel{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groups", errors.New("invalid group id")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	err = groupModel.AddToGroup(myUserModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	UpdateGroupModelContext(c, groupModel.ID)
	serializer := GroupSerializer{c}
	c.JSON(http.StatusAccepted, gin.H{"group": serializer.Response()})
}
