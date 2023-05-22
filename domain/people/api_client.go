package people

type PeopleAPIClient interface {
	GetPeopleFromAPIClientByURLEnding(urlEnding int) (*People, error)
}
