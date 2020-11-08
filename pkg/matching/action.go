package matching

import (
	"errors"
	"log"
	"sync"
)

//ErrCantGetParticipantsNow is retrived when repository returns an error
var ErrCantGetParticipantsNow = errors.New("Can't get participants now")

//MatchingParticipant represents a Participant that match with a Project
type MatchingParticipant struct {
	Name     string
	Distance float32
	Score    float32
}

//Action represents the action of get Participants that matches with a Project
type Action interface {
	GetMatchingParticipantsForProject(p Project) ([]MatchingParticipant, error)
}

type action struct {
	Participants ParticipantRepository
}

func (a *action) GetMatchingParticipantsForProject(p Project) ([]MatchingParticipant, error) {
	wg := sync.WaitGroup{}
	participantsChan := make(chan Participant)
	errChan := make(chan error, 1)

	go func() {
		wg.Wait()
		close(participantsChan)
		close(errChan)
	}()

	for _, city := range p.Cities {
		wg.Add(1)
		go a.getParticipantsPerCity(errChan, participantsChan, city, &wg)
	}

	matchingParticipants := []MatchingParticipant{}

	for p := range participantsChan {
		matchingParticipants = append(matchingParticipants, MatchingParticipant{
			Name: p.Name,
		})
	}

	err := <-errChan
	if err != nil {
		return []MatchingParticipant{}, ErrCantGetParticipantsNow
	}

	return matchingParticipants, nil
}

func (a *action) getParticipantsPerCity(errors chan error, participants chan Participant, city City, wg *sync.WaitGroup) {
	defer wg.Done()
	cityParticipants, err := a.Participants.GetByFormattedAddress(city.FormattedAddress)
	if err != nil {
		log.Println(err)
		errors <- err
		return
	}
	for _, p := range cityParticipants {
		participants <- p
	}
	return
}

func NewMatchingParticipantsAction(repository ParticipantRepository) Action {
	return &action{
		Participants: repository,
	}
}
