package models

import (
	"github.com/VictorKabata/quotes-api/api/auth"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

//Generates user token based on user returned from email and password query
func (login *Auth) SignInUser(db *gorm.DB, email, username, password string) (*Auth, error) {
	user := User{}
	err := db.Debug().Model(&User{}).Where("email=? OR username=?", email, username).Take(&user).Error
	if err != nil {
		return &Auth{}, err
	}

	err = VerifyHashedPassword(password, user.Password)
	if err != nil {
		return &Auth{}, err
	}

	generatedToken, err := auth.CreateToken(user.ID)
	if err != nil {
		return &Auth{}, err
	}

	login.User = user
	login.Token = generatedToken

	return login, nil
}
