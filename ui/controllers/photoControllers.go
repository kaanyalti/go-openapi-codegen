package controllers

import (
	"openapiCodegen/application"

	"github.com/labstack/echo/v4"
)

type PhotoControllers struct {
	photoService application.PhotoService
}

func NewPhotoControllers(photoService application.PhotoService) *PhotoControllers {
	return &PhotoControllers{
		photoService: photoService,
	}
}

func (p *PhotoControllers) CreatePhoto(ctx echo.Context) error {
	return nil
}
