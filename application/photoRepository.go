package application

import (
	domainModels "openapiCodegen/domain/models"
)

type PhotoRepository interface {
	CreatePhoto(domainModels.Photo) error
}
