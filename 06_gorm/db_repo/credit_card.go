package db_repo

import (
	"gorm.io/gorm"
)

type CreditCardService struct {
	db *gorm.DB
}

type CreditCard struct {
	gorm.Model
	UserID      int64   `gorm:"column:user_id"`
	Number      int64   `gorm:"column:number"`
	Balance     float64 `gorm:"column:balance;type:decimal(18,2)"`
	CreditLimit float64 `gorm:"column:credit_limit;type:decimal(18,2)"`
	User        *User   `gorm:"foreignKey:UserID"`
}

func NewCreditCardService(db *gorm.DB) *CreditCardService {
	return &CreditCardService{db: db}
}

func (c *CreditCard) TableName() string {
	return "credit_cards"
}

func (c *CreditCardService) Migrate() error {
	err := c.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&CreditCard{})
	return err
}
