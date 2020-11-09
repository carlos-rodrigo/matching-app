package matching

//ParticipantRepository is an interface where can access to Participants for projects
type ParticipantRepository interface {
	GetByFormattedAddress(address string) ([]Participant, error)
}
