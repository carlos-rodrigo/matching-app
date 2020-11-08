package matching

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockParticipantRepostory struct {
	mock.Mock
}

func (m *mockParticipantRepostory) GetByFormattedAddress(address string) ([]Participant, error) {
	args := m.Called(address)
	return args.Get(0).([]Participant), args.Error(1)
}

func TestMatchingProjectWithParticipants(t *testing.T) {
	distanceService := NewDistanceService()
	city := "New York, NY, USA"
	project := Project{
		Cities: []City{
			City{
				ID:               "ChIJOwg_06VPwokRYv534QaPC8g",
				City:             "New York",
				State:            "NY",
				Country:          "US",
				FormattedAddress: city,
				Location: Location{
					Latitude:  40.7127753,
					Longitude: -74.0059728,
				},
			},
		},
	}

	t.Run("Given a Project, When we don't found participants, Then should return an empty list of participants", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", city).Return([]Participant{}, nil)
		action := NewMatchingParticipantsAction(repository, distanceService)

		participants, _ := action.GetMatchingParticipantsForProject(project)

		assert.Equal(t, 0, len(participants))
	})
	t.Run("Given a Project, When we found participants that match, then should return a list of matching participants", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", city).Return([]Participant{
			Participant{
				Name: "Jefferson",
			},
			Participant{
				Name: "Jillian",
			},
		}, nil)
		action := NewMatchingParticipantsAction(repository, distanceService)

		participants, _ := action.GetMatchingParticipantsForProject(project)

		assert.NotEqual(t, 0, len(participants))
		assert.Equal(t, 2, len(participants))
		for _, participant := range participants {
			assert.NotEqual(t, nil, participant.Name)
			assert.NotEqual(t, nil, participant.Distance)
			assert.NotEqual(t, nil, participant.Score)
		}
		repository.AssertExpectations(t)
	})
	t.Run("Given a Project, When can't access to participants, then should return an user friendly error", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", city).Return([]Participant{}, errors.New("Can't access to storage now"))
		action := NewMatchingParticipantsAction(repository, distanceService)

		_, err := action.GetMatchingParticipantsForProject(project)

		assert.NotNil(t, err)
		assert.Equal(t, "Can't get participants now", err.Error())
		repository.AssertExpectations(t)
	})
	t.Run("Given a Project with a set of cities, When look for participants for the project, Then participants must be located in one of the project cities", func(t *testing.T) {
		project := Project{
			Cities: []City{
				City{
					ID:               "ChIJOwg_06VPwokRYv534QaPC8g",
					City:             "New York",
					State:            "NY",
					Country:          "US",
					FormattedAddress: "New York, NY, USA",
					Location: Location{
						Latitude:  40.7127753,
						Longitude: -74.0059728,
					},
				},
				City{
					ID:               "ChIJ60u11Ni3xokRwVg-jNgU9Yk",
					City:             "Philadelphia",
					State:            "PA",
					Country:          "US",
					FormattedAddress: "Philadelphia, PA, USA",
					Location: Location{
						Latitude:  39.9525839,
						Longitude: -75.1652215,
					},
				},
			},
		}
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", "New York, NY, USA").Return([]Participant{
			Participant{
				Name:             "Jefferson",
				FormattedAddress: "New York, NY, USA",
			},
			Participant{
				Name:             "Jillian",
				FormattedAddress: "New York, NY, USA",
			},
		}, nil)
		repository.On("GetByFormattedAddress", "Philadelphia, PA, USA").Return([]Participant{
			Participant{
				Name:             "Matthew",
				FormattedAddress: "Philadelphia, PA, USA",
			},
		}, nil)
		action := NewMatchingParticipantsAction(repository, distanceService)

		participants, err := action.GetMatchingParticipantsForProject(project)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(participants))
		repository.AssertExpectations(t)

	})

	t.Run("Given a project with specified cities, when match participants for project, then must filter those farther than 100km", func(t *testing.T) {
		project := Project{
			Cities: []City{
				City{
					ID:               "ChIJOwg_06VPwokRYv534QaPC8g",
					City:             "New York",
					State:            "NY",
					Country:          "US",
					FormattedAddress: "New York, NY, USA",
					Location: Location{
						Latitude:  40.7127753,
						Longitude: -74.0059728,
					},
				},
				City{
					ID:               "ChIJ60u11Ni3xokRwVg-jNgU9Yk",
					City:             "Philadelphia",
					State:            "PA",
					Country:          "US",
					FormattedAddress: "Philadelphia, PA, USA",
					Location: Location{
						Latitude:  39.9525839,
						Longitude: -75.1652215,
					},
				},
			},
		}
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", "New York, NY, USA").Return([]Participant{
			Participant{
				Name:             "Jefferson",
				FormattedAddress: "New York, NY, USA",
				Location: Location{
					Latitude:  40.247201,
					Longitude: -74.796316,
				},
			},
			Participant{
				Name:             "Jillian",
				FormattedAddress: "New York, NY, USA",
				Location: Location{
					Latitude:  40.6781784,
					Longitude: -73.9441579,
				},
			},
		}, nil)
		repository.On("GetByFormattedAddress", "Philadelphia, PA, USA").Return([]Participant{
			Participant{
				Name:             "Matthew",
				FormattedAddress: "Philadelphia, PA, USA",
				Location: Location{
					Latitude:  39.9525839,
					Longitude: -75.1652215,
				},
			},
		}, nil)
		action := NewMatchingParticipantsAction(repository, distanceService)

		participants, err := action.GetMatchingParticipantsForProject(project)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(participants))
		repository.AssertExpectations(t)
	})
}
