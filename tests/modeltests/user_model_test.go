package modeltests

import (
	"log"
	"testing"

	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	seededUser := models.User{
		ID:       0,
		Username: "Test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	savedUser, err := seededUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("Error saving user: %v\n", err)
		return
	}

	assert.Equal(t, seededUser.ID, savedUser.ID, "User IDs should match")
	assert.Equal(t, seededUser.Username, savedUser.Username, "Usernames should match")
	assert.Equal(t, seededUser.Email, savedUser.Email, "Emails should match")
}

func TestGetAllUsers(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedMultipleUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.GetAllUsers(server.DB)
	if err != nil {
		log.Fatal(err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestGetUser(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	user, err := userInstance.GetUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, seededUser.ID, user.ID, "User IDs should match")
	assert.Equal(t, seededUser.Username, user.Username, "Usernames should match")
	assert.Equal(t, seededUser.Email, user.Email, "Emails should match")

}

func TestUpdateUser(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	seededUserUpdate := models.User{
		ID:       1,
		Username: "Test Update",
		Email:    "testupdate@gmail.com",
		Password: "passwordupdate",
	}

	updatedUser, err := seededUserUpdate.UpdateUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, seededUserUpdate.ID, updatedUser.ID, "User IDs should match")
	assert.Equal(t, seededUserUpdate.Username, updatedUser.Username, "Usernames should match")
	assert.Equal(t, seededUserUpdate.Email, updatedUser.Email, "Emails should match")

}

func TestDeleteUser(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	isDeleted, err := seededUser.DeleteUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, int(isDeleted), 0)

	allUser, err := userInstance.GetAllUsers(server.DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, len(*allUser), 0)

}
