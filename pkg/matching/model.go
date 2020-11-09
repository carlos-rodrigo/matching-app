package matching

//Project represent a Respondant project
type Project struct {
	Cities                []City   `json:"cities"`
	Genders               string   `json:"genders"`
	ProfessionalIndustry  []string `json:"professionalIndustry"`
	ProfessionalJobTitles []string `json:"professionalJobTitles"`
}

type City struct {
	CityLocation CityLocation `json:"location"`
}

type CityLocation struct {
	ID               string   `json:"id"`
	City             string   `json:"city"`
	State            string   `json:"state"`
	Country          string   `json:"country"`
	FormattedAddress string   `json:"formattedAddress"`
	Location         Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//Participant represent a Respondent or project participant
type Participant struct {
	Name             string
	Gender           string
	FormattedAddress string
	Location         Location
	JobTitle         string
	Industry         []string
}
