package matching

//Project represent a Respondant project
type Project struct {
	Cities                []City
	ProfessionalIndustry  []string
	ProfessionalJobTitles []string
}

type City struct {
	ID               string
	City             string
	State            string
	Country          string
	FormattedAddress string
	Location         Location
}

type Location struct {
	Latitude  float64
	Longitude float64
}

//Participant represent a Respondent or project participant
type Participant struct {
	Name             string
	FormattedAddress string
	Location         Location
	JobTitle         string
	Industry         []string
}
