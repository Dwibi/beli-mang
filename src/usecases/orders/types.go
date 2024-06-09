package ordersusecase

import (
	"github.com/Dwibi/beli-mang/src/entities"
	orderrepository "github.com/Dwibi/beli-mang/src/repositories/order"
	orderitemrepository "github.com/Dwibi/beli-mang/src/repositories/order_items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
)

type sOrdersUseCase struct {
	userRepository      userrepository.IUserRepository
	orderRepository     orderrepository.IOrderRepository
	orderItemRepository orderitemrepository.IOrderItemRepository
}

type IOrdersUseCase interface {
	// Create(m *CreateMerchantParams) (int, int, error)
	// Create(m *CreateMerchantParams) (int, int, error)
	// FindManyParams(m *FindManyParams) (*entities.ItemsResult, int, error)
	// FindManyParams(m *FindManyParams) (*entities.MerchantResult, int, error)
	Create(o CreateOrderParams) (string, int, error)
	FindMany(m *FindManyParams) (*[]entities.ResultListOrderItems, int, error)
}

func New(userRepository userrepository.IUserRepository,
	orderRepository orderrepository.IOrderRepository,
	orderItemRepository orderitemrepository.IOrderItemRepository) IOrdersUseCase {
	return &sOrdersUseCase{
		userRepository:      userRepository,
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
	}
}
