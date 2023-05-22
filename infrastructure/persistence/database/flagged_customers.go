package database

import (
	"github.com/MichaelYoung87/kundbild-public/domain/flaggedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
	"gorm.io/gorm"
	"log"
)

type FlaggedCustomersDBRepository struct {
	db *gorm.DB
}

func NewFlaggedCustomersDBRepository(db *gorm.DB) flaggedcustomers.FlaggedCustomersRepository {
	return &FlaggedCustomersDBRepository{db: db}
}

func (r *FlaggedCustomersDBRepository) SaveFlaggedCustomersToDatabase(flaggedCustomers *flaggedcustomers.FlaggedCustomers) error {
	return r.db.Create(flaggedCustomers).Error
}

func (r *FlaggedCustomersDBRepository) FindFlaggedCustomersInDatabaseByID(id uint) (*flaggedcustomers.FlaggedCustomers, error) {
	var flaggedCustomers flaggedcustomers.FlaggedCustomers
	if err := r.db.Preload("People").Preload("Planets").First(&flaggedCustomers, id).Error; err != nil {
		return nil, err
	}
	return &flaggedCustomers, nil
}

func (r *FlaggedCustomersDBRepository) FindPeopleAndPlanetsInDatabaseByID(peopleID uint, planetsID uint) (*flaggedcustomers.FlaggedCustomers, error) {
	var flaggedCustomers flaggedcustomers.FlaggedCustomers
	if err := r.db.Where("people_id = ? AND planets_id = ?", peopleID, planetsID).First(&flaggedCustomers).Error; err != nil {
		return nil, err
	}
	return &flaggedCustomers, nil
}

func (r *FlaggedCustomersDBRepository) FindPeopleAndPlanetsInDatabaseByURL(peopleURL valueobjects.URL, planetsURL valueobjects.URL) (*flaggedcustomers.FlaggedCustomers, error) {
	var flaggedCustomers flaggedcustomers.FlaggedCustomers
	log.Printf("Searching for FlaggedCustomer with people URL: %s and planets URL: %s", peopleURL.String(), planetsURL.String())
	if err := r.db.Where("people_url = ? AND planets_url = ?", peopleURL.String(), planetsURL.String()).First(&flaggedCustomers).Error; err != nil {
		return nil, err
	}
	return &flaggedCustomers, nil
}
