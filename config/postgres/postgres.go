package postgres

import (
	"os"
	"split_bills/model"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	conn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(conn))

	if err != nil {
		log.Errorf("cant connect to databse %s", err)
	}
	db.AutoMigrate(&model.Users{}, &model.Bills{}, &model.Transactions{})
	return db
}
