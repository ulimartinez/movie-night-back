package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"movie-back/common"
)

type UserModel struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
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
