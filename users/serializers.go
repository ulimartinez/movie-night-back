package users

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
)

type LoginSerializer struct {
	C    *gin.Context
	User UserModel
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (selfr *LoginSerializer) Response() LoginResponse {
	myUserModel := selfr.User
	if myUserModel.ID == 0 {
		myUserModel = selfr.C.MustGet("my_user_model").(UserModel)
	}
	user := LoginResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}

type UserSerializer struct {
	C    *gin.Context
	User UserModel
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (selfr *UserSerializer) Response() UserResponse {
	myUserModel := selfr.User
	if myUserModel.ID == 0 {
		myUserModel = selfr.C.MustGet("my_user_model").(UserModel)
	}
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
	}
	return user
}

type DiscordSerializer struct {
	C	*gin.Context
	Discord DiscordModel
}

type DiscordResponse struct {
	Userid	string	`json:"userid"`
	Token	string	`json:"token"`
}
type DiscordsSerializer struct {
	c *gin.Context
	dcs []DiscordModel
}

func (selfr *DiscordsSerializer) Response() []DiscordResponse {
	response := []DiscordResponse{}
	for _, dc := range selfr.dcs {
		serializer := DiscordSerializer{selfr.c, dc}
		response = append(response, serializer.Response()
	}
	return response
}

func (selfr *DiscordSerializer) Response() DiscordResponse {
	discordModel := selfr.Discord
	discordResponse := DiscordResponse{
		Userid:	discordModel.UserId
		Token:	discordModel.Token
	}
	return discordResponse
}
