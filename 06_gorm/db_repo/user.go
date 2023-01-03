package db_repo

import (
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	UserName     string        `gorm:"column:user_name;type:varchar(1024)"`
	Email        string        `gorm:"column:email;type:varchar(1024)"`
	MobileNumber string        `gorm:"column:mobile_number;type:bigint"`
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
