package itemsusecase

import (
	"github.com/Dwibi/beli-mang/src/entities"
	itemsrepository "github.com/Dwibi/beli-mang/src/repositories/items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
)

type sItemsUseCase struct {
	userRepository  userrepository.IUserRepository
	itemsRepository itemsrepository.IItemsRepository
}

type IItemsUseCase interface {
	// Create(m *CreateMerchantParams) (int, int, error)
	Create(m *CreateMerchantParams) (int, int, error)
	FindManyParams(m *FindManyParams) (*entities.ItemsResult, int, error)
	// FindManyParams(m *FindManyParams) (*entities.MerchantResult, int, error)
}

func New(userRepository userrepository.IUserRepository, itemsRepository itemsrepository.IItemsRepository) IItemsUseCase {
	return &sItemsUseCase{
		userRepository:  userRepository,
		itemsRepository: itemsRepository,
	}
}
