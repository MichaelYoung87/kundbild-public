package planets

type Repository interface {
	FindPlanetsInDatabaseByID(id uint) (*Planets, error)
	FindPlanetsInDatabaseByFullURL(url string) (*Planets, error)
	SavePlanetsToDatabase(planets *Planets) error
}
