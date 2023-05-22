package peopleapi

import (
	"encoding/json"
	"fmt"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"io"
	"net/http"
)

// APIPeople represents people returned from the external API
type APIPeople struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PeopleAPIConfig struct {
	SWAPIBaseURL string
}

type PeopleAPIClient struct {
	Config     PeopleAPIConfig
	HTTPClient *http.Client
}

// NewPeopleAPIClient initializes a new PeopleAPIClient
func NewPeopleAPIClient(SWAPIBaseURL string, httpClient *http.Client) *PeopleAPIClient {
	return &PeopleAPIClient{
		Config: PeopleAPIConfig{
			SWAPIBaseURL: SWAPIBaseURL,
		},
		HTTPClient: httpClient,
	}
}

func (client *PeopleAPIClient) GetPeopleFromAPIClientByURLEnding(urlEnding int) (*people.People, error) {
	// Creates the correct request URL address
	requestURL := fmt.Sprintf("%speople/%d/", client.Config.SWAPIBaseURL, urlEnding)

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
	var apiPeople APIPeople
	err = json.Unmarshal(body, &apiPeople)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Checks if name or url are empty
	if apiPeople.Name == "" || apiPeople.URL == "" {
		return nil, fmt.Errorf("invalid data received from API: empty name or url")
	}

	return convertAPIPeople(apiPeople), nil
}
