package matching

import "errors"

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
	participants, err := a.Participants.GetParticipants()
	if err != nil {
		return []MatchingParticipant{}, ErrCantGetParticipantsNow
	}
	matchingParticipants := make([]MatchingParticipant, len(participants))
	for i, v := range participants {
		matchingParticipants[i] = MatchingParticipant{
			Name: v.Name,
		}
	}

	return matchingParticipants, nil
}

func NewMatchingParticipantsAction(repository ParticipantRepository) Action {
	return &action{
		Participants: repository,
	}
}
