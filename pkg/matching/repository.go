package matching

import "github.com/stretchr/testify/mock"

//ParticipantRepository is an interface where can access to Participants for projects
type ParticipantRepository interface {
	GetParticipants() ([]Participant, error)
}

type mockParticipantRepostory struct {
	mock.Mock
}

func (m *mockParticipantRepostory) GetParticipants() ([]Participant, error) {
	args := m.Called()
	return args.Get(0).([]Participant), args.Error(1)
}
