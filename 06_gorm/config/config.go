package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppCfg struct {
	MYSQLConfig MYSQLConfig
}

type MYSQLConfig struct {
	user            string
	password        string
	host            string
	port            string
	database        string
	dbRetryAttempts int
}

func Init() (*AppCfg, error) {
	var conf AppCfg

	conf.MYSQLConfig = MYSQLConfig{
		user:            os.Getenv("DB_USER"),
		password:        os.Getenv("DB_PASSWORD"),
		host:            os.Getenv("DB_HOST"),
		port:            os.Getenv("DB_PORT"),
		database:        os.Getenv("DB_NAME"),
		dbRetryAttempts: 3,
	}
	return &conf, nil
}

func NewDBConnection(mysqlConf MYSQLConfig) (*gorm.DB, error) {
	var err error
	var connection *gorm.DB
	for i := 0; i < mysqlConf.dbRetryAttempts; i++ {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			mysqlConf.user,
			mysqlConf.password,
			mysqlConf.host,
			// mysqlConf.port,
			mysqlConf.database,
		)
		connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return connection, nil
		}
		time.Sleep(time.Millisecond * 500)
	}
	return nil, err
}
