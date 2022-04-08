package infrastructure

import "portfolioManagement/application"

type gcsPhotoStorage struct {
}

func NewGcsPhotoStorage() application.PhotoStorage {
	return &gcsPhotoStorage{}
}

func (g *gcsPhotoStorage) UploadPhoto() {}
