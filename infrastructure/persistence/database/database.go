package database

import (
	"fmt"
	"github.com/MichaelYoung87/kundbild-public/domain/flaggedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/linkedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(DbUsername, DbPassword, DbHost, DbPort, DbName string) (*gorm.DB, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Europe%%2FStockholm", DbUsername, DbPassword, DbHost, DbPort, DbName)
	connection, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to mysql server: %w", err)
	}
	DB = connection
	err = connection.AutoMigrate(
		&linkedcustomers.LinkedCustomers{},
		&flaggedcustomers.FlaggedCustomers{},
		&people.People{},
		&planets.Planets{},
	)
	if err != nil {
		return nil, fmt.Errorf("could not Auto-Migrate models: %w", err)
	}
	return DB, nil
}
