package uploadsusecase

import (
	"mime/multipart"

	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

type sUploadUseCase struct {
	userRepository userrepository.IUserRepository
	uploader       *manager.Uploader
}

type IUploadUseCase interface {
	//
	Create(userId int, file *multipart.FileHeader) (string, int, error)
}

func New(userRepository userrepository.IUserRepository, uploader *manager.Uploader) IUploadUseCase {
	return &sUploadUseCase{
		userRepository: userRepository,
		uploader:       uploader,
	}
}
