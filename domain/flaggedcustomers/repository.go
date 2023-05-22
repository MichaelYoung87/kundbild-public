package flaggedcustomers

import (
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type FlaggedCustomersRepository interface {
	FindFlaggedCustomersInDatabaseByID(id uint) (*FlaggedCustomers, error)
	FindPeopleAndPlanetsInDatabaseByID(peopleID, planetsID uint) (*FlaggedCustomers, error)
	FindPeopleAndPlanetsInDatabaseByURL(peopleURL, planetsURL valueobjects.URL) (*FlaggedCustomers, error)
	SaveFlaggedCustomersToDatabase(flaggedCustomers *FlaggedCustomers) error
}
