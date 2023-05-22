package peopleapi

import (
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

func convertAPIPeople(apiPeople APIPeople) *people.People {
	return &people.People{
		PeopleName: valueobjects.Name(apiPeople.Name),
		PeopleURL:  valueobjects.URL(apiPeople.URL),
	}
}
