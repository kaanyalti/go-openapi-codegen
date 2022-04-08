package infrastructure

import (
	"portfolioManagement/application"
)

type postgresPhotoRepository struct {
}

func NewPostgresPhotoRepository() application.PhotoRepository {
	return &postgresPhotoRepository{}
}

func (p *postgresPhotoRepository) CreatePhoto() {}
