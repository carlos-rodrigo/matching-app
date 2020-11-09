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
	scoreService := NewScoreService()
	city := "New York, NY, USA"
	projectWithOneCity := Project{
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
	projectWithTwoCities := Project{
		ProfessionalIndustry: []string{
			"Banking",
			"Financial Services",
			"Government Administration",
			"Insurance",
			"Retail",
			"Supermarkets",
			"Automotive",
			"Computer Software",
		},
		ProfessionalJobTitles: []string{
			"Developer",
			"Software Engineer",
			"Software Developer",
			"Programmer",
			"Java Developer",
			"Java/J2EE Developer",
			"Java Full Stack Developer",
			"Java Software Engineer",
			"Java Software Developer",
			"Application Architect",
			"Application Developer",
		},

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

	newYorkPaticipantsWithLessThan100KmDistance := []Participant{
		Participant{
			Name:             "Jefferson",
			Gender:           "male",
			FormattedAddress: "New York, NY, USA",
			Location: Location{
				Latitude:  40.7127753,
				Longitude: -74.0059728,
			},
			JobTitle: "Software Engineer",
		},
		Participant{
			Name:             "Jillian",
			Gender:           "famele",
			FormattedAddress: "New York, NY, USA",
			Location: Location{
				Latitude:  40.6781784,
				Longitude: -73.9441579,
			},
			JobTitle: "Senior Software Engineer",
		},
	}
	phillyParticipantsWithLessThan100KmDistance := []Participant{
		Participant{
			Name:             "Matthew",
			JobTitle:         "Senior Software Engineer",
			FormattedAddress: "Philadelphia, PA, USA",
			Location: Location{
				Latitude:  39.9525839,
				Longitude: -75.1652215,
			},
		},
	}
	twoNewYorkPaticipantsOneWithLessThan100KmDistance := []Participant{
		Participant{
			Name:             "Jefferson",
			FormattedAddress: "New York, NY, USA",
			Location: Location{
				Latitude:  39.9525839,
				Longitude: -75.1652215,
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
	}

	t.Run("Given a Project, When we don't found participants, Then should return an empty list of participants", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", city).Return([]Participant{}, nil)
		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		participants, _ := action.GetMatchingParticipantsForProject(projectWithOneCity)

		assert.Equal(t, 0, len(participants))
	})
	t.Run("Given a Project, When we found participants that match, then should return a list of matching participants", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", city).Return(newYorkPaticipantsWithLessThan100KmDistance, nil)
		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		participants, _ := action.GetMatchingParticipantsForProject(projectWithOneCity)

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
		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		_, err := action.GetMatchingParticipantsForProject(projectWithOneCity)

		assert.NotNil(t, err)
		assert.Equal(t, "Can't get participants now", err.Error())
		repository.AssertExpectations(t)
	})
	t.Run("Given a Project with a set of cities, When look for participants for the project, Then participants must be located in one of the project cities", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", "New York, NY, USA").Return(newYorkPaticipantsWithLessThan100KmDistance, nil)
		repository.On("GetByFormattedAddress", "Philadelphia, PA, USA").Return(phillyParticipantsWithLessThan100KmDistance, nil)
		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		participants, err := action.GetMatchingParticipantsForProject(projectWithTwoCities)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(participants))
		repository.AssertExpectations(t)

	})

	t.Run("Given a Project with a set of cities, When look for participants for the project, Then participants must be located in less than 100km", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", "New York, NY, USA").Return(twoNewYorkPaticipantsOneWithLessThan100KmDistance, nil)
		repository.On("GetByFormattedAddress", "Philadelphia, PA, USA").Return(phillyParticipantsWithLessThan100KmDistance, nil)
		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		participants, err := action.GetMatchingParticipantsForProject(projectWithTwoCities)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(participants))
		repository.AssertExpectations(t)

	})

	t.Run("Given a project with a defined industry and job titles, When participants are found, Then they must be sorted by match score in desendent order", func(t *testing.T) {
		repository := new(mockParticipantRepostory)
		repository.On("GetByFormattedAddress", "New York, NY, USA").Return(newYorkPaticipantsWithLessThan100KmDistance, nil)
		repository.On("GetByFormattedAddress", "Philadelphia, PA, USA").Return(phillyParticipantsWithLessThan100KmDistance, nil)

		action := NewMatchingParticipantsAction(repository, distanceService, scoreService)

		participants, err := action.GetMatchingParticipantsForProject(projectWithTwoCities)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(participants))
		assert.True(t, participants[0].Score >= participants[1].Score, "%+v p1 %+v p2", participants[0], participants[2])
		assert.True(t, participants[1].Score >= participants[2].Score, "%+v p1 %+v p2", participants[1], participants[2])
		repository.AssertExpectations(t)
	})
}
