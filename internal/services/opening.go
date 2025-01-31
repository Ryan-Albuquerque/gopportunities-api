package services

import (
	"github.com/ryan-albuquerque/gopportunities-api/internal/models"
	"github.com/ryan-albuquerque/gopportunities-api/internal/repositories"
)

func CreateOpening(openingRequest models.CreateOpeningRequest) (uint, error) {
	opening := &models.Opening{
		Role:      openingRequest.Role,
		Company:   openingRequest.Company,
		Location:  openingRequest.Location,
		Link:      openingRequest.Link,
		Remote:    openingRequest.Remote,
		Salary:    openingRequest.Salary,
		IsActived: true,
	}

	err := repositories.CreateOpening(opening)
	if err != nil {
		return 0, err
	}

	return opening.ID, nil
}

func ListOpenings() ([]models.Opening, error) {
	return repositories.ListOpenings()
}
