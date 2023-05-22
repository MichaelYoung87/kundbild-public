package services

import (
	"errors"
	"github.com/MichaelYoung87/kundbild-public/domain/linkedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"gorm.io/gorm"
	"time"
)

type LinkedCustomersService struct {
	linkedCustomersRepo linkedcustomers.LinkedCustomersRepository
	peopleRepo          people.Repository
	planetsRepo         planets.Repository
}

func NewLinkedCustomersService(linkedCustomersRepo linkedcustomers.LinkedCustomersRepository, peopleRepo people.Repository, planetsRepo planets.Repository) *LinkedCustomersService {
	return &LinkedCustomersService{
		linkedCustomersRepo: linkedCustomersRepo,
		peopleRepo:          peopleRepo,
		planetsRepo:         planetsRepo,
	}
}

func (lcs *LinkedCustomersService) SaveLinkedCustomersToDatabase(people *people.People, planets *planets.Planets) (*linkedcustomers.LinkedCustomers, bool, error) {

	// Try to find the existing object of people by URL
	existingPeople, err := lcs.peopleRepo.FindPeopleInDatabaseByFullURL(people.PeopleURL.String())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// If an object of people is found, assign it to the people variable
		people = existingPeople
	}

	// Try to find the existing object of planets by URL
	existingPlanets, err := lcs.planetsRepo.FindPlanetsInDatabaseByFullURL(planets.PlanetsURL.String())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// If an object of planets is found, assign it to the planets variable
		planets = existingPlanets
	}

	// Check if the combination of people and planets already exists in linked customers
	existingLinkedCustomers, err := lcs.linkedCustomersRepo.FindPeopleAndPlanetsInDatabaseByURL(people.PeopleURL, planets.PlanetsURL)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// Return existing linked customers with flag set to false
		return existingLinkedCustomers, false, nil
	}

	// If an object of people and planets were not found earlier, save them now
	if existingPeople == nil {
		err = lcs.peopleRepo.SavePeopleToDatabase(people)
		if err != nil {
			return nil, false, err
		}
	}
	if existingPlanets == nil {
		err = lcs.planetsRepo.SavePlanetsToDatabase(planets)
		if err != nil {
			return nil, false, err
		}
	}

	createdAtTime := time.Now()

	linkedCustomers := &linkedcustomers.LinkedCustomers{
		People:        people,
		PeopleURL:     people.PeopleURL,
		Planets:       planets,
		PlanetsURL:    planets.PlanetsURL,
		CreatedAtTime: createdAtTime,
	}

	err = lcs.linkedCustomersRepo.SaveLinkedCustomersToDatabase(linkedCustomers)
	if err != nil {
		return nil, false, err
	}

	return linkedCustomers, true, nil
}

func (lcs *LinkedCustomersService) FindPeopleInDatabaseByID(id uint) (*people.People, error) {
	return lcs.peopleRepo.FindPeopleInDatabaseByID(id)
}

func (lcs *LinkedCustomersService) FindPlanetsInDatabaseByID(id uint) (*planets.Planets, error) {
	return lcs.planetsRepo.FindPlanetsInDatabaseByID(id)
}
