package matching

//Project represent a Respondant project
type Project struct {
	Cities []City
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
	Latitude  float32
	Longitude float32
}

//Participant represent a Respondent or project participant
type Participant struct {
	Name             string
	FormattedAddress string
	Location         Location
}
