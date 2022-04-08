package application

type PhotoService interface {
	CreatePhoto()
}

type photoService struct {
	photoRepository PhotoRepository
	photoStorage    PhotoStorage
}

func NewPhotoService(photoRepository PhotoRepository, photoStorage PhotoStorage) PhotoService {
	return &photoService{
		photoRepository: photoRepository,
		photoStorage:    photoStorage,
	}
}

func (p *photoService) CreatePhoto() {}
