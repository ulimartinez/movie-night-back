package groups

import (
	"github.com/gin-gonic/gin"
	"movie-back/users"
)

type GroupSerializer struct {
	c     *gin.Context
	group GroupModel
}

type GroupResponse struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func (selfr *GroupSerializer) Response() GroupResponse {
	myGroupModel := selfr.group
	if myGroupModel.ID == 0 {
		myGroupModel = selfr.c.MustGet("my_group_model").(GroupModel)
	}
	group := GroupResponse{
		Name: myGroupModel.Name,
		ID:   myGroupModel.ID,
	}
	return group
}

type GroupsSerializer struct {
	c      *gin.Context
	Groups []GroupModel
}

func (s *GroupsSerializer) Response() []GroupResponse {
	response := []GroupResponse{}
	for _, group := range s.Groups {
		serializer := GroupSerializer{s.c, group}
		response = append(response, serializer.Response())
	}
	return response
}

type UsersSerializer struct {
	c     *gin.Context
	Users []users.UserModel
}

func (s *UsersSerializer) Response() []users.UserResponse {
	response := []users.UserResponse{}
	for _, user := range s.Users {
		serializer := users.UserSerializer{s.c, user}
		response = append(response, serializer.Response())
	}
	return response
}
