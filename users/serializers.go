package users

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
)

type UserSerializer struct {
	C    *gin.Context
	User UserModel
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (selfr *UserSerializer) Response() UserResponse {
	myUserModel := selfr.User
	if myUserModel.ID == 0 {
		myUserModel = selfr.C.MustGet("my_user_model").(UserModel)
	}
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
