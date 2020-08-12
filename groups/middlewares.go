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
				if len(groups) <= 0 {
					UpdateGroupModelContext(c, 0)
				} else {
					UpdateGroupModelContext(c, groups[0].ID)
				}
			}
		}
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
