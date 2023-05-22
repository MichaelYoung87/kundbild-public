package linkedcustomers

import (
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type LinkedCustomersRepository interface {
	FindLinkedCustomersInDatabaseByID(id uint) (*LinkedCustomers, error)
	FindPeopleAndPlanetsInDatabaseByID(peopleID uint, planetsID uint) (*LinkedCustomers, error)
	FindPeopleAndPlanetsInDatabaseByURL(peopleURL valueobjects.URL, planetsURL valueobjects.URL) (*LinkedCustomers, error)
	SaveLinkedCustomersToDatabase(linkedCustomers *LinkedCustomers) error
}
