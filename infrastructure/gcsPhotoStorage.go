package infrastructure

import "portfolioManagement/services"

type gcsPhotoStorage struct {
}

func NewGcsPhotoStorage() services.PhotoStorage {
	return &gcsPhotoStorage{}
}

func (g *gcsPhotoStorage) UploadPhoto() {}
