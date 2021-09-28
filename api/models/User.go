package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment;unique" json:"id"`
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;unique" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//Encrypts password - Returns hashed password string
func (user *User) HashPassword(password string) (string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordByte), err
}

//Compares hashed password and unencryped password - Returns error
func (user *User) VerifyHashedPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

//Hashes password before creating new user
func (user *User) HashPasswordBeforeSave() error {
	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return nil
}

//Resets values to default values
func (user *User) Prepare() {
	user.ID = 0
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

//Validates input
func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email format")
		}

		if user.Email == "" {
			return errors.New("Email required")
		}
		if user.Password == "" {
			return errors.New("Password required")
		}
	}

	return nil
}

//Saves new user to database
func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

//Returns all users saved in database
func (user *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}

	err := db.Debug().Model(&User{}).Find(&user).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

//Returns a specific user querried from the database
func (user *User) GetUser(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

//Updates specific users records in the database
func (user *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).Updates(User{Username: user.Username, Email: user.Email, Password: user.Password}).Error
	if err != nil {
		return &User{}, err
	}

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

//Deletes specific user record from the database
func (user *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	deleteTx := db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Delete(&user)
	if deleteTx != nil {
		return 0, deleteTx.Error
	}

	return deleteTx.RowsAffected, nil
}

//Generates user token based on user returned from email and password query
func (user *User) SignInUser(db *gorm.DB, email, password string) (string, error) {
	err := db.Model(&User{}).Where("email=? AND password=?", email, password).Take(&user).Error
	if err != nil {
		return "", err
	}

	// generatedToken,err:=auth.CreateToken(user.ID)
	// if err!=nil{
	// 	return "", err
	// }

	return "", nil
}
