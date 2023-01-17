package controllers

import "github.com/labstack/echo/v4"

type PhotoController struct {}

func NewPhotoController() PhotoController {
  return PhotoController{}
}

func (c PhotoController) CreatePhoto(ctx echo.Context) error {
}
  
func (c PhotoController) GetPhotos(ctx echo.Context) error {
}
  
