package infrastructure

import (
	"portfolioManagement/application"
	domainModels "portfolioManagement/domain/models"
	dbModels "portfolioManagement/infrastructure/models"

	"gorm.io/gorm"
)

type postgresPhotoRepository struct {
	db *gorm.DB
}

func NewPostgresPhotoRepository(db *gorm.DB) application.PhotoRepository {
	return &postgresPhotoRepository{
		db: db,
	}
}

func (p *postgresPhotoRepository) CreatePhoto(photo domainModels.Photo) error {
	ph := dbModels.Photo{
		UUID: photo.UUID,
		Name: photo.Name,
		Size: photo.Size,
	}

	result := p.db.Create(&ph)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
