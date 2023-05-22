package services

import (
	"errors"
	"github.com/MichaelYoung87/kundbild-public/domain/flaggedcustomers"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"gorm.io/gorm"
	"time"
)

type FlaggedCustomersService struct {
	flaggedCustomersRepo flaggedcustomers.FlaggedCustomersRepository
	peopleRepo           people.Repository
	planetsRepo          planets.Repository
}

func NewFlaggedCustomersService(flaggedCustomersRepo flaggedcustomers.FlaggedCustomersRepository, peopleRepo people.Repository, planetsRepo planets.Repository) *FlaggedCustomersService {
	return &FlaggedCustomersService{
		flaggedCustomersRepo: flaggedCustomersRepo,
		peopleRepo:           peopleRepo,
		planetsRepo:          planetsRepo,
	}
}

func (fcs *FlaggedCustomersService) SaveFlaggedCustomersToDatabase(people *people.People, planets *planets.Planets) (*flaggedcustomers.FlaggedCustomers, bool, error) {

	// Try to find the existing people in the database by the full URL address
	existingPeople, err := fcs.peopleRepo.FindPeopleInDatabaseByFullURL(people.PeopleURL.String())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// If an object of people is found, assign it to the people variable
		people = existingPeople
	}

	// Try to find the existing planets in the database by the full URL address
	existingPlanets, err := fcs.planetsRepo.FindPlanetsInDatabaseByFullURL(planets.PlanetsURL.String())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// If an object of planets is found, assign it to the planets variable
		planets = existingPlanets
	}

	// Check if the combination of people and planets already exists in flagged customers
	existingFlaggedCustomers, err := fcs.flaggedCustomersRepo.FindPeopleAndPlanetsInDatabaseByURL(people.PeopleURL, planets.PlanetsURL)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, err
		}
	} else {
		// Return existing flagged customers with flag set to false
		return existingFlaggedCustomers, false, nil
	}

	// If an object of people and planets were not found earlier, save them now
	if existingPeople == nil {
		err = fcs.peopleRepo.SavePeopleToDatabase(people)
		if err != nil {
			return nil, false, err
		}
	}
	if existingPlanets == nil {
		err = fcs.planetsRepo.SavePlanetsToDatabase(planets)
		if err != nil {
			return nil, false, err
		}
	}

	createdAtTime := time.Now()

	flaggedCustomers := &flaggedcustomers.FlaggedCustomers{
		People:        people,
		PeopleURL:     people.PeopleURL,
		Planets:       planets,
		PlanetsURL:    planets.PlanetsURL,
		CreatedAtTime: createdAtTime,
	}

	err = fcs.flaggedCustomersRepo.SaveFlaggedCustomersToDatabase(flaggedCustomers)
	if err != nil {
		return nil, false, err
	}

	return flaggedCustomers, true, nil
}

// Hard coded People and Planets that will be a match to be able to get flagged
func (fcs *FlaggedCustomersService) CheckMatch(people *people.People, planets *planets.Planets) bool {
	matchingCombinations := map[string]string{
		"Luke Skywalker": "Hoth",     //	"Luke Skywalker"	= people/1/		"Hoth"		= planets/4/
		"Leia Organa":    "Tatooine", //	"Leia Organa"		= people/5/		"Tatooine"	= planets/1/
		"Darth Vader":    "Dagobah",  // 	"Darth Vader"		= people/4/		"Dagobah"	= planets/5/
		"C-3PO":          "Yavin IV", //	"C-3PO"				= people/2/		"Yavin IV"	= planets/3/
		"R2-D2":          "Alderaan", //	"R2-D2"				= people/3/		"Alderaan"	= planets/2/
		"Owen Lars":      "Bespin",   //	"Owen Lars"			= people/6/		"Bespin"	= planets/6/
	}

	matchingPlanet, ok := matchingCombinations[people.PeopleName.String()]
	if ok && matchingPlanet == planets.PlanetsName.String() {
		return true
	}

	return false
}
