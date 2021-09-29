package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Quote struct {
	ID        uint32    `gorm:"primary_key;auto_increment;unique" json:"id"`
	Statement string    `gorm:"size:255;not null;unique" json:"quote"`
	Sayer     string    `gorm:"size:100;not null" json:"sayer"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (quote *Quote) Prepare() {
	quote.ID = 0
	quote.Statement = html.EscapeString(strings.TrimSpace(quote.Statement))
	quote.Sayer = html.EscapeString(strings.TrimSpace(quote.Sayer))
	quote.User = User{}
	quote.CreatedAt = time.Now()
	quote.UpdatedAt = time.Now()
}

func (quote *Quote) ValidateInput(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if quote.Statement == "" {
			return errors.New("Statement required")
		}
		if quote.Sayer == "" {
			return errors.New("Sayer required")
		}
	}

	return nil
}

func (quote *Quote) SaveQuote(db *gorm.DB) (*Quote, error) {
	err := db.Debug().Model(&Quote{}).Create(&quote).Error
	if err != nil {
		return &Quote{}, err
	}

	if quote.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id=?", quote.UserID).Take(&quote.User).Error
		if err != nil {
			return &Quote{}, err
		}
	}

	return quote, nil
}

func (quote *Quote) GetAllQuotes(db *gorm.DB) (*[]Quote, error) {
	quotes := []Quote{}

	err := db.Debug().Model(&Quote{}).Find(&quotes).Error
	if err != nil {
		return &[]Quote{}, err
	}

	if len(quotes) > 0 {
		for key, _ := range quotes {
			err := db.Debug().Model(&User{}).Where("id=?", quotes[key].ID).Take(&quotes[key].User).Error
			if err != nil {
				return &[]Quote{}, err
			}
		}
	}

	return &quotes, nil
}

func (quote *Quote) GetQuote(db *gorm.DB, qid uint32) (*Quote, error) {
	err := db.Debug().Model(&Quote{}).Where("id=?", qid).Take(&Quote{}).Error
	if err != nil {
		return &Quote{}, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Quote{}, errors.New("Quote not found")
	}

	if quote.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id=?", quote.UserID).Take(&quote.User).Error
		if err != nil {
			return &Quote{}, err
		}
	}

	return quote, nil
}

func (quote *Quote) UpdateQuote(db *gorm.DB) (*Quote, error) {
	err := db.Debug().Model(&Quote{}).Where("id=?", quote.ID).Updates(Quote{Statement: quote.Statement, Sayer: quote.Sayer, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Quote{}, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Quote{}, errors.New("Quote not found")
	}

	if quote.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id=?", quote.UserID).Take(&quote.User).Error
		if err != nil {
			return &Quote{}, err
		}
	}

	return quote, nil
}

func (quote *Quote) DeleteQuote(db *gorm.DB, qid uint32) (int64, error) {
	deleteTx := db.Debug().Model(&Quote{}).Where("id=?", qid).Take(&Quote{}).Delete(&Quote{})
	if deleteTx.Error != nil {
		return 0, deleteTx.Error
	}

	if gorm.IsRecordNotFoundError(deleteTx.Error) {
		return 0, errors.New("Quote not found")
	}

	return deleteTx.RowsAffected, nil
}
