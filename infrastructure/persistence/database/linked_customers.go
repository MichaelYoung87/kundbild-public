package database

import (
	"github.com/MichaelYoung87/kundbild-public/domain/linkedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
	"gorm.io/gorm"
	"log"
)

type LinkedCustomersDBRepository struct {
	db *gorm.DB
}

func NewLinkedCustomersDBRepository(db *gorm.DB) linkedcustomers.LinkedCustomersRepository {
	return &LinkedCustomersDBRepository{db: db}
}

func (r *LinkedCustomersDBRepository) SaveLinkedCustomersToDatabase(linkedCustomers *linkedcustomers.LinkedCustomers) error {
	return r.db.Create(linkedCustomers).Error
}

func (r *LinkedCustomersDBRepository) FindLinkedCustomersInDatabaseByID(id uint) (*linkedcustomers.LinkedCustomers, error) {
	var linkedCustomers linkedcustomers.LinkedCustomers
	if err := r.db.Preload("People").Preload("Planets").First(&linkedCustomers, id).Error; err != nil {
		return nil, err
	}
	return &linkedCustomers, nil
}

func (r *LinkedCustomersDBRepository) FindPeopleAndPlanetsInDatabaseByID(peopleID uint, planetsID uint) (*linkedcustomers.LinkedCustomers, error) {
	var linkedCustomers linkedcustomers.LinkedCustomers
	if err := r.db.Where("people_id = ? AND planets_id = ?", peopleID, planetsID).First(&linkedCustomers).Error; err != nil {
		return nil, err
	}
	return &linkedCustomers, nil
}

func (r *LinkedCustomersDBRepository) FindPeopleAndPlanetsInDatabaseByURL(peopleURL valueobjects.URL, planetsURL valueobjects.URL) (*linkedcustomers.LinkedCustomers, error) {
	var linkedCustomers linkedcustomers.LinkedCustomers
	log.Printf("Searching for LinkedCustomer with people URL: %s and planets URL: %s", peopleURL.String(), planetsURL.String())
	if err := r.db.Where("people_url = ? AND planets_url = ?", peopleURL.String(), planetsURL.String()).First(&linkedCustomers).Error; err != nil {
		return nil, err
	}
	return &linkedCustomers, nil
}
