package matching

import (
	"errors"
	"log"
	"sort"
	"sync"
)

const maxDistance = 100.00

//ErrCantGetParticipantsNow is retrived when repository returns an error
var ErrCantGetParticipantsNow = errors.New("Can't get participants now")

//MatchingParticipant represents a Participant that match with a Project
type MatchingParticipant struct {
	Name       string
	Distance   float64
	Score      float64
	LocationID string
}

type byScore []MatchingParticipant

func (s byScore) Len() int {
	return len(s)
}

func (s byScore) Less(i, j int) bool {
	return s[j].Score < s[i].Score
}

func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//DistanceParticipant represents a tuple Participant and Distance to a location
type DistanceParticipant struct {
	Participant Participant
	Distance    float64
	LocationID  string
}

//Action represents the action of get Participants that matches with a Project
type Action interface {
	GetMatchingParticipantsForProject(p Project) ([]MatchingParticipant, error)
}

type action struct {
	Participants ParticipantRepository
	Distance     DistanceService
	Score        ScoreService
}

func (a *action) GetMatchingParticipantsForProject(project Project) ([]MatchingParticipant, error) {
	wg := sync.WaitGroup{}
	participantsChan := make(chan DistanceParticipant)
	errChan := make(chan error, 1)

	go func() {
		wg.Wait()
		close(participantsChan)
		close(errChan)
	}()

	for _, city := range project.Cities {
		wg.Add(1)
		go a.getParticipantsPerCity(errChan, participantsChan, city, &wg)
	}

	matchingParticipants := []MatchingParticipant{}

	for distanceParticipant := range participantsChan {
		matchingParticipants = append(matchingParticipants, MatchingParticipant{
			Name:       distanceParticipant.Participant.Name,
			Score:      a.Score.GetMatchingScore(project, distanceParticipant.Participant),
			Distance:   distanceParticipant.Distance,
			LocationID: distanceParticipant.LocationID,
		})
	}

	err := <-errChan
	if err != nil {
		return []MatchingParticipant{}, ErrCantGetParticipantsNow
	}

	sort.Sort(byScore(matchingParticipants))

	return matchingParticipants, nil
}

func (a *action) getParticipantsPerCity(errors chan error, participants chan DistanceParticipant, city City, wg *sync.WaitGroup) {
	defer wg.Done()
	cityParticipants, err := a.Participants.GetByFormattedAddress(city.FormattedAddress)
	if err != nil {
		log.Println(err)
		errors <- err
		return
	}
	for _, p := range cityParticipants {
		distance := a.Distance.GetDistanceBetweenLocations(p.Location, city.Location)
		if distance <= maxDistance {
			participants <- DistanceParticipant{
				Participant: p,
				Distance:    distance,
				LocationID:  city.ID,
			}
		}
	}
	return
}

//NewMatchingParticipantsAction returns an Action to get the matching participants from a project
func NewMatchingParticipantsAction(repository ParticipantRepository, distance DistanceService, score ScoreService) Action {
	return &action{
		Participants: repository,
		Distance:     distance,
		Score:        score,
	}
}
