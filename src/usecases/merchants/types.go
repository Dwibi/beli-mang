package merchantusecase

import (
	"github.com/Dwibi/beli-mang/src/entities"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
)

type sMerchantUseCase struct {
	userRepository     userrepository.IUserRepository
	merchantRepository merchantrepository.IMerchantRepository
}

type IMerchantUseCase interface {
	Create(m *CreateMerchantParams) (int, int, error)
	FindManyParams(m *FindManyParams) (*entities.MerchantResult, int, error)
}

func New(userRepository userrepository.IUserRepository, merchantRepository merchantrepository.IMerchantRepository) IMerchantUseCase {
	return &sMerchantUseCase{
		userRepository:     userRepository,
		merchantRepository: merchantRepository,
	}
}
