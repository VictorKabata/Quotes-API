package modeltests

import (
	"log"
	"testing"

	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/stretchr/testify/assert"
)

func TestSaveQuote(t *testing.T) {
	dropUsersandQuotesTables()

	mockUser := models.User{
		ID:       1,
		Username: "Mary Doe",
		Email:    "mary@gmail.com",
		Password: "password",
	}

	mockQuote := models.Quote{
		ID:        1,
		Statement: "Test statement",
		Sayer:     "Tester",
		UserID:    mockUser.ID,
		User:      mockUser,
	}

	savedQuote, err := mockQuote.SaveQuote(server.DB)
	if err != nil {
		log.Fatalf("Error creating new quote: %v\n", err)
		return
	}

	assert.Equal(t, mockQuote.ID, savedQuote.ID, "Quote IDs should match")
	assert.Equal(t, mockQuote.Statement, savedQuote.Statement, "Statements should match")
	assert.Equal(t, mockQuote.Sayer, savedQuote.Sayer, "Sayer should match")
	assert.Equal(t, mockQuote.UserID, savedQuote.UserID, "User IDs match")
	assert.Equal(t, mockQuote.User, savedQuote.User, "Users should match")
}

func TestGetAllQuotes(t *testing.T) {
	err := dropUsersandQuotesTables()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = seedMultipleQuotes()
	if err != nil {
		log.Fatal(err)
		return
	}

	quotes, err := quoteInstance.GetAllQuotes(server.DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, len(*quotes), 2)
}

func TestGetQuote(t *testing.T) {
	err := dropUsersandQuotesTables()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}

	seededQuote, err := seedOneQuote()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}

	foundQuote, err := quoteInstance.GetQuote(server.DB, seededQuote.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}

	assert.Equal(t, foundQuote.ID, seededQuote.ID)
	assert.Equal(t, foundQuote.Statement, seededQuote.Statement)
	assert.Equal(t, foundQuote.Sayer, seededQuote.Sayer)
	assert.Equal(t, foundQuote.UserID, seededQuote.UserID)
}

func TestUpdateQuote(t *testing.T) {
	err := dropUsersandQuotesTables()
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = seedOneQuote()
	if err != nil {
		log.Fatal(err)
		return
	}

	user := models.User{
		ID:       1,
		Username: "Mary Doe",
		Email:    "mary@gmail.com",
		Password: "password",
	}

	seededQuoteUpdate := models.Quote{
		ID:        1,
		Statement: "Test statement update",
		Sayer:     "Tester update",
		UserID:    user.ID,
		User:      user,
	}

	updatedQuote, err := seededQuoteUpdate.UpdateQuote(server.DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, seededQuoteUpdate.Statement, updatedQuote.Statement)
	assert.Equal(t, seededQuoteUpdate.Sayer, updatedQuote.Sayer)
	assert.Equal(t, seededQuoteUpdate.UserID, updatedQuote.UserID)
	assert.Equal(t, seededQuoteUpdate.User, updatedQuote.User)
}

func TestDeleteQuote(t *testing.T) {
	err := dropUsersandQuotesTables()
	if err != nil {
		log.Fatal(err)
		return
	}

	seededQuote, err := seedOneQuote()
	if err != nil {
		log.Fatal(err)
		return
	}

	isDeleted, err := quoteInstance.DeleteQuote(server.DB, seededQuote.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, int(isDeleted), 1)
}
