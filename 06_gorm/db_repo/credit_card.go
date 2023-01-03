package db_repo

import (
	"time"

	"gorm.io/gorm"
)

type CreditCardService struct {
	db *gorm.DB
}

type CreditCard struct {
	ID int64     `gorm:"primary_key;autoIncrement;column:credit_card_id"`
	UserID       int64     `gorm:"column:user_id"`
	Number       int64     `gorm:"column:number"`
	Balance      float64   `gorm:"column:balance;type:decimal(18,2)"`
	CreditLimit  float64   `gorm:"column:credit_limit;type:decimal(18,2)"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime"`
	CreatedBy    int64     `gorm:"column:created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime"`
	UpdatedBy    int64     `gorm:"column:updated_by"`
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime"`
	DeletedBy    int64     `gorm:"column:deleted_by"`
	User         *User     `gorm:"foreignKey:UserID"`
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
