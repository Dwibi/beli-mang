package estimatesrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sEstimatesRepository struct {
	DB *sql.DB
}

type IEstimatesRepository interface {
	// FindMany(merchantId int, filters *entities.SearchItemsParams) (*entities.ItemsResult, error)
	// Create(merchantId int, m *entities.CreateItemsParams) (int, error)
	Create(userId int, totalPrice, estimatedDeliveryTime float64) (*entities.ResultEstimate, error)
	Update(totalPrice float64) (*entities.ResultEstimate, error)
}

func New(db *sql.DB) IEstimatesRepository {
	return &sEstimatesRepository{DB: db}
}
