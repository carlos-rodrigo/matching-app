package matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreService(t *testing.T) {
	t.Run("Given a project with a professional industries, When matching score is evaluated, Then must recieve one point for every industry match", func(t *testing.T) {
		service := NewScoreService()
		project := Project{
			ProfessionalIndustry: []string{
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
	t.Run("Given a project with JobTitles, When matching score is evaluated, Then add one point to score if the participant job title is a full match with project expected job titles", func(t *testing.T) {
		service := NewScoreService()
		project := Project{
			ProfessionalJobTitles: []string{
				"Developer",
				"Software Engineer",
				"Software Developer",
				"Programmer",
				"Java Developer",
				"Java/J2EE Developer",
				"Java Full Stack Developer",
				"Java Software Engineer",
				"Java Software Developer",
				"Application Architect",
				"Application Developer",
			},
		}
		participant := Participant{
			JobTitle: "Software Engineer",
		}

		score := service.GetMatchingScore(project, participant)

		assert.Equal(t, 1.0, score)
	})

	t.Run("Given a project with JobTitles, When matching score is evaluated, Then add half point to score if the participant job title is a full match with project expected job titles and contains a high seniority indicatior", func(t *testing.T) {
		service := NewScoreService()
		project := Project{
			ProfessionalJobTitles: []string{
				"Developer",
				"Software Engineer",
				"Software Developer",
				"Programmer",
				"Java Developer",
				"Java/J2EE Developer",
				"Java Full Stack Developer",
				"Java Software Engineer",
				"Java Software Developer",
				"Application Architect",
				"Application Developer",
			},
		}
		participant := Participant{
			JobTitle: "Senior Software Engineer",
		}

		score := service.GetMatchingScore(project, participant)

		assert.Equal(t, 1.5, score)
	})

}
