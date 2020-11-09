package matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceService(t *testing.T) {
	t.Run("Given two locations (York, Bristol), When calculate distance between them, Then distance should be 296.71", func(t *testing.T) {
		york := Location{
			Longitude: 1.0803,
			Latitude:  53.9583,
		}
		bristol := Location{
			Longitude: 2.5833,
			Latitude:  51.4500,
		}
		service := NewDistanceService()

		distance := service.GetDistanceBetweenLocations(york, bristol)

		assert.Equal(t, 296.7075963774831, distance)
	})
	t.Run("Given the same locations (York), When calculate distance between them, Then distance should be 0", func(t *testing.T) {
		york := Location{
			Longitude: 1.0803,
			Latitude:  53.9583,
		}
		york2 := Location{
			Longitude: 1.0803,
			Latitude:  53.9583,
		}
		service := NewDistanceService()

		distance := service.GetDistanceBetweenLocations(york, york2)

		assert.Equal(t, 0.0, distance)
	})

}
