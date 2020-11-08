package matching

type DistanceService interface {
	GetDistanceBetweenLocations(p1, p2 Location) float32
}

type distanceService struct {
}

func (s *distanceService) GetDistanceBetweenLocations(p1, p2 Location) float32 {
	panic("Not implemented distance method.")
}

func NewDistanceService() DistanceService {
	return &distanceService{}
}
