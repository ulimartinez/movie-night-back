package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := LoginSerializer{c, userModelValidator.userModel}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&UserModel{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered email or invalid password")))
		return
	}
	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered email or invalid password")))
		return
	}
	UpdateContextUserModel(c, userModel.ID)
	serializer := LoginSerializer{c, UserModel{ID: 0}}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
