package matching

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ErrCantGetParticipantsNow = errors.New("Can't get participants now")

type Project struct{}

type Participant struct {
	Name string
}

type MatchingParticipant struct {
	Name     string
	Distance float32
	Score    float32
}

type MatchingParticipantsAction interface {
	GetMatchingParticipantsForProject(p Project) ([]MatchingParticipant, error)
}

type action struct {
	Participants ParticipantRepository
}

func (a *action) GetMatchingParticipantsForProject(p Project) ([]MatchingParticipant, error) {
	participants, err := a.Participants.GetParticipants()
	if err != nil {
		return []MatchingParticipant{}, ErrCantGetParticipantsNow
	}
	matchingParticipants := make([]MatchingParticipant, len(participants))
	for i, v := range participants {
		matchingParticipants[i] = MatchingParticipant{
			Name: v.Name,
		}
	}

	return matchingParticipants, nil
}

func NewMatchingParticipantsAction(repository ParticipantRepository) MatchingParticipantsAction {
	return &action{
		Participants: repository,
	}
}

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

func TestMatchingProjectWithParticipants(t *testing.T) {
	t.Run("Given a Project, When we don't found participants, Then should return an empty list of participants", func(t *testing.T) {
		project := Project{}
		repository := new(mockParticipantRepostory)
		repository.On("GetParticipants").Return([]Participant{}, nil)
		action := NewMatchingParticipantsAction(repository)

		participants, _ := action.GetMatchingParticipantsForProject(project)

		assert.Equal(t, 0, len(participants))
	})
	t.Run("Given a Project, When we found participants that match, then should return a list of matching participants", func(t *testing.T) {
		project := Project{}
		repository := new(mockParticipantRepostory)
		repository.On("GetParticipants").Return([]Participant{
			Participant{
				Name: "Jefferson",
			},
			Participant{
				Name: "Jillian",
			},
		}, nil)
		action := NewMatchingParticipantsAction(repository)

		participants, _ := action.GetMatchingParticipantsForProject(project)

		assert.NotEqual(t, 0, len(participants))
		assert.Equal(t, 2, len(participants))
		for _, participant := range participants {
			assert.NotEqual(t, nil, participant.Name)
			assert.NotEqual(t, nil, participant.Distance)
			assert.NotEqual(t, nil, participant.Score)
		}
	})
	t.Run("Given a Project, When can't access to participants, then should return an user friendly error", func(t *testing.T) {
		project := Project{}
		repository := new(mockParticipantRepostory)
		repository.On("GetParticipants").Return([]Participant{}, errors.New("Can't access to storage now"))
		action := NewMatchingParticipantsAction(repository)

		_, err := action.GetMatchingParticipantsForProject(project)

		assert.NotNil(t, err)
		assert.Equal(t, "Can't get participants now", err.Error())
	})

}
