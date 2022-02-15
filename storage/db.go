package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB(host, user, dbname, password, port string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&UserProfile{})

	return db, nil
}
