package planets

type PlanetsAPIClient interface {
	GetPlanetsFromAPIClientByURLEnding(urlEnding int) (*Planets, error)
}
