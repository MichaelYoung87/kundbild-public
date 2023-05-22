package database

import (
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"gorm.io/gorm"
)

type PeopleDBRepository struct {
	db *gorm.DB
}

func NewPeopleDBRepository(db *gorm.DB) *PeopleDBRepository {
	return &PeopleDBRepository{db: db}
}

func (r *PeopleDBRepository) SavePeopleToDatabase(peopleEntity *people.People) error {
	return r.db.Where(&people.People{PeopleURL: peopleEntity.PeopleURL}).FirstOrCreate(peopleEntity).Error
}

func (r *PeopleDBRepository) FindPeopleInDatabaseByFullURL(url string) (*people.People, error) {
	var peopleEntity people.People
	if err := r.db.Where("people_url = ?", url).First(&peopleEntity).Error; err != nil {
		return nil, err
	}
	return &peopleEntity, nil
}

func (r *PeopleDBRepository) FindPeopleInDatabaseByID(id uint) (*people.People, error) {
	var peopleEntity people.People
	if err := r.db.First(&peopleEntity, id).Error; err != nil {
		return nil, err
	}
	return &peopleEntity, nil
}
