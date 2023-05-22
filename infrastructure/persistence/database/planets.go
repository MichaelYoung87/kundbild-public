package database

import (
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"gorm.io/gorm"
)

type PlanetsDBRepository struct {
	db *gorm.DB
}

func NewPlanetsDBRepository(db *gorm.DB) planets.Repository {
	return &PlanetsDBRepository{db: db}
}

func (r *PlanetsDBRepository) SavePlanetsToDatabase(planetsEntity *planets.Planets) error {
	return r.db.Where(&planets.Planets{PlanetsURL: planetsEntity.PlanetsURL}).FirstOrCreate(planetsEntity).Error
}

func (r *PlanetsDBRepository) FindPlanetsInDatabaseByFullURL(url string) (*planets.Planets, error) {
	var planetsEntity planets.Planets
	if err := r.db.Where("planets_url = ?", url).First(&planetsEntity).Error; err != nil {
		return nil, err
	}
	return &planetsEntity, nil
}

func (r *PlanetsDBRepository) FindPlanetsInDatabaseByID(id uint) (*planets.Planets, error) {
	var planetsEntity planets.Planets
	if err := r.db.First(&planetsEntity, id).Error; err != nil {
		return nil, err
	}
	return &planetsEntity, nil
}
