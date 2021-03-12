package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/ravener/discord-oauth2"
	"io/ioutil"
	"movie-back/common"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
	router.OPTIONS("/login", preflight)
	router.GET("/discord", DiscordFetch)
	router.POST("/discord", DiscordAdd)
	router.GET("/auth", DiscordExchange)
}

func DiscordExchange(c *gin.Context) {
	type DiscordUser struct {
		Id			string
		Username		string
		Avatar			string
		Discriminator		string
		PublicFlags		string	`json::public_flags"`
		Flags			int
		Locale			string
		MfaEnabled		bool	`json:"mfa_enabled"`
		PremiumType		int	`json:"premium_type"`
	}
	var user DiscordUser
	conf := &oauth2.Config{
		RedirectURL: "http://localhost:3000/api/users/auth",
		ClientID: "761072595419398164",
		ClientSecret: "Tiidf-T6wD4h8n6rrgyZzdZvW5SAoIBC",
		Scopes: []string{discord.ScopeIdentify},
		Endpoint: discord.Endpoint,
	}
	movie_night_token := c.Request.URL.Query().Get("state")
	token, err := conf.Exchange(oauth2.NoContext, c.Request.URL.Query().Get("code"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("discord", err))
	}
	res, err := conf.Client(oauth2.NoContext, token).Get("https://discordapp.com/api/v7/users/@me")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("discord", err))
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(body), &user)


	if err := SaveOne(&DiscordModel{UserId: user.Id, Token: movie_night_token}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := LoginSerializer{c, userModelValidator.userModel}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&UserModel{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered email or invalid password")))
		return
	}
	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered email or invalid password")))
		return
	}
	UpdateContextUserModel(c, userModel.ID)
	serializer := LoginSerializer{c, UserModel{ID: 0}}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func DiscordFetch(c *gin.Context) {
	discordModel, err := FindDiscordUsers(DiscordModel{})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("discord", errors.New("no users found")))
		return
	}
	serializer := DiscordsSerializer{c, discordModel}
	c.JSON(http.StatusOK, gin.H{"discord": serializer.Response()})
}

func DiscordAdd(c *gin.Context) {
	discordModelValidator := NewDiscordModelValidator()
	if err := discordModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := SaveOne(&discordModelValidator.discordModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := DiscordSerializer{c, discordModelValidator.discordModel}
	c.JSON(http.StatusCreated, gin.H{"discord": serializer.Response()})
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
