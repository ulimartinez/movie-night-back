package groups

import "github.com/gin-gonic/gin"

type GroupSerializer struct {
	c *gin.Context
}

type GroupResponse struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func (selfr *GroupSerializer) Response() GroupResponse {
	myGroupModel := selfr.c.MustGet("my_group_model").(GroupModel)
	group := GroupResponse{
		Name: myGroupModel.Name,
		ID:   myGroupModel.ID,
	}
	return group
}
