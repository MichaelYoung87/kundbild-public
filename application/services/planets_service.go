package services

import (
	"fmt"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type PlanetsService struct {
	planetsRepo      planets.Repository
	planetsAPIClient planets.PlanetsAPIClient
}

func NewPlanetsService(planetsRepo planets.Repository, planetsAPIClient planets.PlanetsAPIClient) *PlanetsService {
	return &PlanetsService{
		planetsRepo:      planetsRepo,
		planetsAPIClient: planetsAPIClient,
	}
}

func (s *PlanetsService) GetPlanetsFromAPIClientByURLEnding(urlEnding int) (*planets.Planets, error) {
	p, err := s.planetsAPIClient.GetPlanetsFromAPIClientByURLEnding(urlEnding)
	if err != nil {
		return nil, err
	}

	planetsEntity := &planets.Planets{
		PlanetsID:   p.PlanetsID,
		PlanetsName: valueobjects.Name(p.PlanetsName),
		PlanetsURL:  valueobjects.URL(p.PlanetsURL),
	}
	return planetsEntity, nil
}

func (s *PlanetsService) SavePlanetsToDatabase(planetsEntity *planets.Planets) error {
	fmt.Println("Storing an object of Planets:", planetsEntity.PlanetsName)
	err := s.planetsRepo.SavePlanetsToDatabase(planetsEntity)
	if err != nil {
		fmt.Println("Error when saving Planet:", err)
		return err
	}
	return nil
}

func (s *PlanetsService) FindPlanetsInDatabaseByID(id uint) (*planets.Planets, error) {
	return s.planetsRepo.FindPlanetsInDatabaseByID(id)
}
