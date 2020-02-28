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
	router.GET("/list", UserGroups)
	router.GET("/users/:grid", GroupUsers)
	router.POST("/set/:id", SetGroup)
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
	serializer := GroupSerializer{c, groupModel}
	c.JSON(http.StatusAccepted, gin.H{"group": serializer.Response()})
}

func UserGroups(c *gin.Context) {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	groups, err := GetGroups(myUserModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := GroupsSerializer{c, groups}
	c.JSON(http.StatusOK, gin.H{"groups": serializer.Response()})
}

func GroupUsers(c *gin.Context) {
	groupId, err := strconv.ParseUint(c.Param("grid"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("conversion error"))
	}
	myGroupModel, err := FindOneGroup(&GroupModel{ID: uint(groupId)})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	usersList, err := GetUsers(myGroupModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	serializer := UsersSerializer{c, usersList}
	c.JSON(http.StatusOK, gin.H{"users": serializer.Response()})
}

func SetGroup(c *gin.Context) {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	groupId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("conversion", err))
	}
	if groupId == 0 {
		groupModel, err := GetGroups(myUserModel)
		if err != nil {
			c.JSON(http.StatusNotFound, common.NewError("database", err))
		}
		groupId = uint64(groupModel[0].ID)
	}
	UpdateGroupModelContext(c, uint(groupId))
	serializer := GroupSerializer{c, GroupModel{ID: 0}}
	c.JSON(http.StatusOK, gin.H{"group": serializer.Response()})
}
