package application

import (
	domainModels "portfolioManagement/domain/models"
)

type PhotoRepository interface {
	CreatePhoto(domainModels.Photo) error
}
