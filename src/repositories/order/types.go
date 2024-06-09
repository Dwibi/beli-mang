package orderrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sOrderRepository struct {
	DB *sql.DB
}

type IOrderRepository interface {
	Create(reqUserId int, CalculatedEstimateId string) (int, error)
	FindMany(userId int, filters *entities.SearchOrderParams) (*[]entities.ResultListOrderItems, error)
}

func New(db *sql.DB) IOrderRepository {
	return &sOrderRepository{DB: db}
}
