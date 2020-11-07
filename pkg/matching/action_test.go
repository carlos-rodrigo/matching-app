package matching

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
