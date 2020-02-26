package main

import (
	"github.com/gin-gonic/contrib/static"
	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/users"
	"net/http"
)

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	v1.Use(users.AuthMiddleware(true))
	testAuth := r.Group("/api/ping")
	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}
