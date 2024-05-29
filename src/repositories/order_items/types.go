package orderitemrepository

import (
	"database/sql"
)

type sOrderItemRepository struct {
	DB *sql.DB
}

type IOrderItemRepository interface {
	Create(reqUserId int, CalculatedEstimateId string) error
}

func New(db *sql.DB) IOrderItemRepository {
	return &sOrderItemRepository{DB: db}
}
