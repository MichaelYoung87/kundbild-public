package planets

import (
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type Planets struct {
	PlanetsID   uint              `json:"planets_ID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	PlanetsName valueobjects.Name `json:"planets_name"`
	PlanetsURL  valueobjects.URL  `json:"planets_URL"`
}
