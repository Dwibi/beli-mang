package estimatesusecase

import (
	"github.com/Dwibi/beli-mang/src/entities"
	estimateitemsrepository "github.com/Dwibi/beli-mang/src/repositories/estimate_items"
	estimatesrepository "github.com/Dwibi/beli-mang/src/repositories/estimates"
	itemsrepository "github.com/Dwibi/beli-mang/src/repositories/items"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
)

type sEstimatesUseCase struct {
	userRepository          userrepository.IUserRepository
	estimatesRepository     estimatesrepository.IEstimatesRepository
	merchantrepository      merchantrepository.IMerchantRepository
	itemsrepository         itemsrepository.IItemsRepository
	estimateitemsrepository estimateitemsrepository.IEstimateItemsRepository
}

type IEstimatesUseCase interface {
	Create(p CreateEstimateParams) (*entities.ResultEstimate, int, error)
}

func New(userRepository userrepository.IUserRepository, merchantrepository merchantrepository.IMerchantRepository, itemsrepository itemsrepository.IItemsRepository, estimatesRepository estimatesrepository.IEstimatesRepository, estimateitemsrepository estimateitemsrepository.IEstimateItemsRepository) IEstimatesUseCase {
	return &sEstimatesUseCase{
		userRepository:          userRepository,
		estimatesRepository:     estimatesRepository,
		itemsrepository:         itemsrepository,
		merchantrepository:      merchantrepository,
		estimateitemsrepository: estimateitemsrepository,
	}
}
