package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type CommonsError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidatorError(err error) CommonsError {
	res := CommonsError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return res
}

func NewError(key string, err error) CommonsError {
	res := CommonsError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

const NBSecretPassword = "A very secure string that is a password for signing ! OK!"
const NBRandomPassword = "This would be a test password ok!!"

func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Year * 24).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token

}
