package people

type Repository interface {
	FindPeopleInDatabaseByID(id uint) (*People, error)
	FindPeopleInDatabaseByFullURL(url string) (*People, error)
	SavePeopleToDatabase(people *People) error
}
