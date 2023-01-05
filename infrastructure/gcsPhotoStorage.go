package infrastructure

import "openapiCodegen/application"

type gcsPhotoStorage struct {
}

func NewGcsPhotoStorage() application.PhotoStorage {
	return &gcsPhotoStorage{}
}

func (g *gcsPhotoStorage) UploadPhoto() {}
