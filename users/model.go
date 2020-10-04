package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"movie-back/common"
	"time"
)

type UserModel struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
	GroupID      uint
	LastVote     time.Time `time_format:"2020-01-01" gorm:"default:2020-01-01"`
}

type DiscordModel struct {
	Id	uint	`gorm:"primary_key"`
	UserId	string	`gorm:"column:userid"`
	Token	string `gorm:"column:token"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&DiscordModel{})
}
func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func FindDiscordUsers(data interface{}) ([]DiscordModel, error) {
	db := common.GetDB()
	var models []DiscordModel
	err := db.Find(&models).Error
	return models, err
}

func UpdateVoted(data interface{}) error {
	db := common.GetDB()
	return db.Model(data).Update(&UserModel{LastVote: time.Now()}).Error
}
