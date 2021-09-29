package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/VictorKabata/quotes-api/api/controllers"
	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var quoteInstance = models.Quote{}
var err error

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting enviromental variable: %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

//Make a connection to the database
func Database() {
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")

	DBUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	server.DB, err = gorm.Open("postgres", DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Successfully connected to database")
	}
}

//Clear the Users table before every test
func dropUsersTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Println("Dropped users table")
	return nil
}

//Create/Mock one user for testing
func seedOneUser() (models.User, error) {
	dropUsersTable()

	user := models.User{
		Username: "Mary Doe",
		Email:    "mary@gmail.com",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Failed to seed one user: %v", err)
	}

	return user, nil
}

//Create/Mock multiple users for testing
func seedMultipleUsers() ([]models.User, error) {
	dropUsersTable()

	users := []models.User{
		{
			Username: "Mary Doe",
			Email:    "mary@gmail.com",
			Password: "password",
		},

		{
			Username: "John Doe",
			Email:    "john@gmail.com",
			Password: "mementomori",
		},
	}

	for key, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[key]).Error
		if err != nil {
			return []models.User{}, err
		}
	}

	return users, nil
}

//Drop User and Quotes tables before every test
func dropUsersandQuotesTables() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Quote{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Quote{}).Error
	if err != nil {
		return err
	}

	log.Println("Dropped users and quotes table")
	return nil
}

//Creates/mock a single quote
func seedOneQuote() (models.Quote, error) {
	dropUsersandQuotesTables()

	seededUser, _ := seedOneUser()

	quote := models.Quote{
		ID:        1,
		Statement: "Test statement",
		Sayer:     "Tester",
		UserID:    seededUser.ID,
		User:      seededUser,
	}

	err := server.DB.Model(&models.Quote{}).Create(&quote).Error
	if err != nil {
		log.Fatalf("Error seeding quote: %v\n", err)
		return models.Quote{}, nil
	}

	return quote, nil
}

//Create/mocks multiple quotes
func seedMultipleQuotes() error {
	dropUsersandQuotesTables()

	seededUsers, _ := seedMultipleUsers()

	quotes := []models.Quote{
		{
			Statement: "Test statement",
			Sayer:     "Tester",
		},

		{
			Statement: "Test statement 2",
			Sayer:     "Tester 2",
		},
	}

	for key, _ := range seededUsers {
		quotes[key].ID = seededUsers[key].ID
		quotes[key].UserID = seededUsers[key].ID
		quotes[key].User = seededUsers[key]

		err := server.DB.Model(&models.Quote{}).Create(&quotes[key]).Error
		if err != nil {
			log.Fatalf("Error seeding multiple quotes: %v\n", err)
			return err
		}
	}

	return nil
}
