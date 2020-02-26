package groups

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
)

func UpdateGroupModelContext(c *gin.Context, id uint) {
	var myGroupModel GroupModel
	if id != 0 {
		db := common.GetDB()
		db.First(&myGroupModel, id)
	}
	c.Set("my_group_id", id)
	c.Set("my_group_model", myGroupModel)
}
