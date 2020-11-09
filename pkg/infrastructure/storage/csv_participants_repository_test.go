package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvParticipantsRepository(t *testing.T) {
	csvPath := "respondents_data_test.csv"
	t.Run("Given a new initialization, When repository is created, Then csv file fullfil the repository", func(t *testing.T) {
		repository := NewCsvParticipantsRepository(csvPath)

		assert.NotEqual(t, 0, len(repository.Participants))
	})
}
