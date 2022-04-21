package models

import (
	"github.com/google/uuid"
)

type Photo struct {
	UUID uuid.UUID
	Name string
	Size int
}

func NewPhoto(name string, size int) *Photo {
	return &Photo{
		UUID: uuid.New(),
		Name: name,
		Size: size,
	}
}
