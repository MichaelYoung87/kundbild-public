package people

import (
	"github.com/MichaelYoung87/kundbild-public/domain/valueobjects"
)

type People struct {
	PeopleID   uint              `json:"people_ID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	PeopleName valueobjects.Name `json:"people_name"`
	PeopleURL  valueobjects.URL  `json:"people_URL"`
}

// TableName overrides the default table name 'peoples'
func (People) TableName() string {
	return "people"
}
