package flaggedcustomers

import (
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
	"time"
)

type FlaggedCustomers struct {
	FlaggedCustomersID uint `json:"flaggedCustomersID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	PeopleID           uint `json:"people_ID"`
	People             *people.People
	PeopleURL          valueobjects.URL `json:"people_URL" gorm:"type:varchar(256);uniqueIndex:idx_people_url"`
	PlanetsID          uint             `json:"planets_ID"`
	Planets            *planets.Planets
	PlanetsURL         valueobjects.URL `json:"planets_URL" gorm:"type:varchar(256);uniqueIndex:idx_planets_url"`
	CreatedAtTime      time.Time        `json:"createdAtTime"`
}
