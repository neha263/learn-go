package db_repo

import (
	"time"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	ID       int64         `gorm:"primary_key;autoIncrement;column:user_id"`
	UserName     string        `gorm:"column:user_name;type:varchar(1024)"`
	Email        string        `gorm:"column:email;type:varchar(1024)"`
	MobileNumber string        `gorm:"column:mobile_number;type:bigint"`
	CreatedAt    time.Time     `gorm:"column:created_at;type:datetime"`
	CreatedBy    int64         `gorm:"column:created_by"`
	UpdatedAt    time.Time     `gorm:"column:updated_at;type:datetime"`
	UpdatedBy    int64         `gorm:"column:updated_by"`
	DeletedAt    time.Time     `gorm:"column:deleted_at;type:datetime"`
	DeletedBy    int64         `gorm:"column:deleted_by"`
	CreditCards  []*CreditCard `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (c *User) TableName() string {
	return "users"
}

func (c *UserService) Migrate() error {
	err := c.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	return err
}
