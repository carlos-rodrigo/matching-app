package matching

import "strings"

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
	for _, jobTitle := range projectExpectedJobsTitles {
		if strings.ToLower(jobTitle) == strings.ToLower(participantJobTitle) {
			score++
		}
	}

	return score

}

func NewScoreService() ScoreService {
	return &scoreService{}
}
