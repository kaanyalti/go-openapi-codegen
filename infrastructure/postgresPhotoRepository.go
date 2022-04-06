package infrastructure

import (
	"portfolioManagement/services"

	"gorm.io/gorm"
)

type postgresPhotoRepository struct {
}

func NewPostgresPhotoRepository(dbCon *gorm.DB) services.PhotoRepository {
	return &postgresPhotoRepository{}
}

func (p *postgresPhotoRepository) CreatePhoto() {}
