package matching

import "math"

const earthRadius = float64(6371)

type DistanceService interface {
	GetDistanceBetweenLocations(p1, p2 Location) float64
}

type distanceService struct {
}

func (s *distanceService) GetDistanceBetweenLocations(from, to Location) float64 {
	var deltaLat = (to.Latitude - from.Latitude) * (math.Pi / 180)
	var deltaLon = (to.Longitude - from.Longitude) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(from.Latitude*(math.Pi/180))*math.Cos(to.Latitude*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func NewDistanceService() DistanceService {
	return &distanceService{}
}
