package matching

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ScoreService interface {
	GetMatchingScore(project Project, participant Participant) float64
}

type scoreService struct {
}

func (s *scoreService) GetMatchingScore(project Project, participant Participant) float64 {
	score := 0.0
	for _, industry := range participant.Industry {
		for _, projectIndustry := range project.ProfesionalIndustry {
			if strings.ToLower(industry) == strings.ToLower(projectIndustry) {
				score++
			}
		}
	}
	return score
}

func NewScoreService() ScoreService {
	return &scoreService{}
}

func TestScoreService(t *testing.T) {
	t.Run("Given a project with a professional industries, When matching score is evaluated, Then must recieve one point for every industry match", func(t *testing.T) {
		service := NewScoreService()
		project := Project{
			ProfesionalIndustry: []string{
				"Banking",
				"Financial Services",
				"Government Administration",
				"Insurance",
				"Retail",
				"Supermarkets",
				"Automotive",
				"Computer Software",
			},
		}
		participant := Participant{
			Industry: []string{
				"Information Technology and Services",
				"Banking",
				"Computer Software",
				"Computer Hardware",
				"Financial Services",
			},
		}

		score := service.GetMatchingScore(project, participant)

		assert.Equal(t, 3.0, score)
	})
}
