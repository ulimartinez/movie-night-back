package groups

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
)

type GroupValidator struct {
	Group struct {
		Name string `form:"name" json:"name" binding:"required"`
	} `json:"group"`
	groupModel GroupModel `json:"-"`
}

func (self *GroupValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.groupModel.Name = self.Group.Name
	return nil
}
func NewGroupValidator() GroupValidator {
	groupValidator := GroupValidator{}
	return groupValidator
}

type UpdateValidator struct {
	User struct {
		Email string `form:"email" json:"email" binding:"required,email"`
	} `json:"user"`
	Group struct {
		ID uint `form:"id" json:"id" binding:"required"`
	} `json:"group"`
	userModel  users.UserModel `json:"-"`
	groupModel GroupModel      `json:"-"`
}

func (selfr *UpdateValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, selfr)
	if err != nil {
		return err
	}
	selfr.groupModel.ID = selfr.Group.ID
	selfr.userModel.Email = selfr.User.Email
	return nil
}
func NewUpdateValidator() UpdateValidator {
	updateValidator := UpdateValidator{}
	return updateValidator
}
