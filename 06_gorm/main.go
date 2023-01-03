package main

import (
	"gorm_demo/config"
	"gorm_demo/db_repo"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type MysqlRepoServices struct {
	userRepoService       *db_repo.UserService
	creditCardRepoService *db_repo.CreditCardService
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(false)

	conf, err := config.Init()
	if err != nil {
		log.Fatalf("unable to load config details : %v", err)
	}
	log.Info("config details set")
	db, err := config.NewDBConnection(conf.MYSQLConfig)
	if err != nil {
		log.Fatalf("error occured while initiating database connection : %v", err)
	}
	log.Info("db connection successful")

	RegisterRepoServices(db)
	log.Info("db tables created successfully")
}

func RegisterRepoServices(db *gorm.DB) {
	var repos MysqlRepoServices
	repos.userRepoService = db_repo.NewUserService(db)
	if err := repos.userRepoService.Migrate(); err != nil {
		log.Fatal(err)
	}
	repos.creditCardRepoService = db_repo.NewCreditCardService(db)
	if err := repos.creditCardRepoService.Migrate(); err != nil {
		log.Fatal(err)
	}
}
