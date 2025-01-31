package repositories

import (
	"fmt"

	"github.com/ryan-albuquerque/gopportunities-api/internal/config"
	"github.com/ryan-albuquerque/gopportunities-api/internal/models"
)

// CreateOpening is a function that creates a new opening

func CreateOpening(opening *models.Opening) error {
	db := config.GetDB()
	if err := db.Create(opening).Error; err != nil {
		return err
	}
	return nil
}

func ListOpenings() ([]models.Opening, error) {
	db := config.GetDB()

	var openings []models.Opening

	if err := db.Where("is_actived = ?", true).Find(&openings).Error; err != nil {
		return nil, err
	}

	return openings, nil
}

func GetOpeningByID(id uint) (models.Opening, error) {
	db := config.GetDB()
	if db == nil {
		return models.Opening{}, fmt.Errorf("database connection is not initialized")
	}

	var opening models.Opening

	if err := db.First(&opening, id).Error; err != nil {
		return models.Opening{}, err
	}

	return opening, nil
}

	func UpdateOpening(opening models.Opening) error {
	db := config.GetDB()

	return db.Save(&opening).Error
}
