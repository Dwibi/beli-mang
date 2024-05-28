package estimateitemsrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sEstimateItemsRepository struct {
	DB *sql.DB
}

type IEstimateItemsRepository interface {
	FindTotalPrice(calculatedEstimateId int) (float64, error)
	Create(calculatedEstimateId int, orders []entities.Orders) error
}

func New(db *sql.DB) IEstimateItemsRepository {
	return &sEstimateItemsRepository{DB: db}
}
