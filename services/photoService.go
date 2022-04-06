package services

type PhotoService interface {
	CreatePhoto()
}

type photoService struct {
}

func NewPhotoService() PhotoService {
	return &photoService{}
}

func (p *photoService) CreatePhoto() {}
