package groups

import (
	"github.com/gin-gonic/gin"
	"movie-back/common"
	"movie-back/users"
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

func GroupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("my_group_id")
		if exists == false {
			myUserModel := c.MustGet("my_user_model").(users.UserModel)
			groups, err := GetGroups(myUserModel)
			if err == nil {
				UpdateGroupModelContext(c, groups[0].ID)
			}
		}
	}
}
