package planetsapi

import (
	"encoding/json"
	"fmt"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"io"
	"net/http"
)

type APIPlanets struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PlanetsAPIConfig struct {
	SWAPIBaseURL string
}

type PlanetsAPIClient struct {
	Config     PlanetsAPIConfig
	HTTPClient *http.Client
}

// NewPlanetsAPIClient initializes a new PlanetsAPIClient
func NewPlanetsAPIClient(SWAPIBaseURL string, httpClient *http.Client) *PlanetsAPIClient {
	return &PlanetsAPIClient{
		Config: PlanetsAPIConfig{
			SWAPIBaseURL: SWAPIBaseURL,
		},
		HTTPClient: httpClient,
	}
}

func (client *PlanetsAPIClient) GetPlanetsFromAPIClientByURLEnding(urlEnding int) (*planets.Planets, error) {
	// Creates the correct request URL address
	requestURL := fmt.Sprintf("%splanets/%d/", client.Config.SWAPIBaseURL, urlEnding)

	// HTTP GET request
	response, err := client.HTTPClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer response.Body.Close()

	// Checks if response.StatusCode == OK
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Reads the reponse.Body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parses the JSON result
	var apiPlanets APIPlanets
	err = json.Unmarshal(body, &apiPlanets)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Checks if name or url are empty
	if apiPlanets.Name == "" || apiPlanets.URL == "" {
		return nil, fmt.Errorf("invalid data received from API: empty name or url")
	}

	return convertAPIPlanets(apiPlanets), nil
}
