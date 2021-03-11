package main

import (
	"github.com/gin-gonic/contrib/static"
	gin "github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	"movie-back/common"
	"movie-back/groups"
	"movie-back/movies"
	"movie-back/nights"
	"movie-back/users"
	"net/http"
	"time"
)

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.New()

	v1 := r.Group("/api")
	v1.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "user-agent, Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	v1.OPTIONS("/", preflight)
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	v1.Use(users.AuthMiddleware(true))
	v1.Use(groups.GroupMiddleware())
	groups.GroupCreate(v1.Group("/groups"))
	movies.MovieRegister(v1.Group("/movies"))
	nights.NightsRegister(v1.Group("/nights"))
	r.Run(":3000")
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	groups.AutoMigrate()
	movies.Automigrate()
	nights.AutoMigrate()
}
