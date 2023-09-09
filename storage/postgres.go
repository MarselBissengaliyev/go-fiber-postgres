package storage

import (
	"fmt"

	"github.com/MarselBissengaliyev/go-fiber-postgres/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DB_Host,
		config.DB_User,
		config.DB_Password,
		config.DB_Name,
		config.Port,
		config.DB_SSLMode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	
	return db, nil
}
