package matching

import (
	"strings"
)

var highSeniorityIndicators = []string{
	"Sr ",
	"Sr. ",
	"Senior ",
	"Expert ",
	"Principal ",
	"Staff ",
	"Specialist ",
	"Head ",
}

//ScoreService manage all relative calculation to matching score
type ScoreService interface {
	GetMatchingScore(project Project, participant Participant) float64
}

type scoreService struct {
}

func (s *scoreService) GetMatchingScore(project Project, participant Participant) float64 {
	score := evalIndustriesScore(participant.Industry, project.ProfessionalIndustry)
	score += evalJobTitleScore(participant.JobTitle, project.ProfessionalJobTitles)

	return score
}

func evalIndustriesScore(participantIndustries []string, projectIndustries []string) float64 {
	score := 0.0
	for _, industry := range participantIndustries {
		for _, projectIndustry := range projectIndustries {
			if strings.ToLower(industry) == strings.ToLower(projectIndustry) {
				score++
			}
		}
	}

	return score

}

func evalJobTitleScore(participantJobTitle string, projectExpectedJobsTitles []string) float64 {
	score := 0.0
	loweredParticipantJobTitle := strings.ToLower(participantJobTitle)
	for _, jobTitle := range projectExpectedJobsTitles {
		loweredJobTitle := strings.ToLower(jobTitle)
		if loweredJobTitle == loweredParticipantJobTitle {
			score++
		} else {
			if strings.Contains(loweredParticipantJobTitle, loweredJobTitle) {
				score++
				for _, seniority := range highSeniorityIndicators {
					if strings.Contains(loweredParticipantJobTitle, strings.ToLower(seniority)) {
						score += 0.5
					}
				}
			}
		}
	}
	return score
}

//NewScoreService returns a new ScoreService
func NewScoreService() ScoreService {
	return &scoreService{}
}
