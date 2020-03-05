package nights

import (
	"github.com/gin-gonic/gin"
	"movie-back/users"
)

type NightSerializer struct {
	c     *gin.Context
	Night NightModel
}

type NightResponse struct {
	ID       uint `json:"id"`
	GroupID  uint `json:"group_id"`
	Host     users.UserResponse
	Location string `json:"location"`
}

func (selfr *NightSerializer) Response() NightResponse {
	nightModel := selfr.Night
	userSerializer := users.UserSerializer{selfr.c, nightModel.Host}
	response := NightResponse{
		ID:       nightModel.ID,
		GroupID:  nightModel.GroupID,
		Host:     userSerializer.Response(),
		Location: nightModel.Location,
	}
	return response
}

type NightsSerializer struct {
	c      *gin.Context
	Nights []NightModel
}

func (selfr *NightsSerializer) Response() []NightResponse {
	response := []NightResponse{}
	models := selfr.Nights
	for _, night := range models {
		serializer := NightSerializer{selfr.c, night}
		response = append(response, serializer.Response())
	}
	return response
}
