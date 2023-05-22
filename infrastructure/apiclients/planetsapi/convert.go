package planetsapi

import (
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

func convertAPIPlanets(apiPlanets APIPlanets) *planets.Planets {
	return &planets.Planets{
		PlanetsName: valueobjects.Name(apiPlanets.Name),
		PlanetsURL:  valueobjects.URL(apiPlanets.URL),
	}
}
