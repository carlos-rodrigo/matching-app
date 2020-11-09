package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvParticipantsRepository(t *testing.T) {
	csvPath := "respondents_data_test.csv"
	repository := NewCsvParticipantsRepository(csvPath)

	t.Run("Given a new initialization, When repository is created, Then csv file fullfil the repository", func(t *testing.T) {
		assert.NotEqual(t, 0, len(repository.Participants))
	})
	t.Run("Given a formattedAddres, When want to retrieve participants, Then must only return participants that match with the formattedAddress", func(t *testing.T) {
		participants, err := repository.GetByFormattedAddress("New York, NY, USA")

		assert.Nil(t, err)
		assert.NotEqual(t, 50, len(participants))
	})
}
