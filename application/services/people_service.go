package services

import (
	"fmt"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type PeopleService struct {
	peopleRepo      people.Repository
	peopleAPIClient people.PeopleAPIClient
}

func NewPeopleService(peopleRepo people.Repository, peopleAPIClient people.PeopleAPIClient) *PeopleService {
	return &PeopleService{
		peopleRepo:      peopleRepo,
		peopleAPIClient: peopleAPIClient,
	}
}

func (s *PeopleService) GetPeopleFromAPIClientByURLEnding(urlEnding int) (*people.People, error) {
	p, err := s.peopleAPIClient.GetPeopleFromAPIClientByURLEnding(urlEnding)
	if err != nil {
		return nil, err
	}

	peopleEntity := &people.People{
		PeopleID:   p.PeopleID,
		PeopleName: valueobjects.Name(p.PeopleName),
		PeopleURL:  valueobjects.URL(p.PeopleURL),
	}
	return peopleEntity, nil
}

func (s *PeopleService) SavePeopleToDatabase(peopleEntity *people.People) error {
	fmt.Println("Storing an object of People:", peopleEntity.PeopleName)
	err := s.peopleRepo.SavePeopleToDatabase(peopleEntity)
	if err != nil {
		fmt.Println("Error when saving an object of People:", err)
		return err
	}
	return nil
}

func (s *PeopleService) FindPeopleInDatabaseByID(id uint) (*people.People, error) {
	return s.peopleRepo.FindPeopleInDatabaseByID(id)
}
