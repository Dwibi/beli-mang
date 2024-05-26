package itemsrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sItemsRepository struct {
	DB *sql.DB
}

type IItemsRepository interface {
	FindMany(filters *entities.SearchItemsParams) (*entities.ItemsResult, error)
	Create(m *entities.CreateItemsParams) (int, error)
}

func New(db *sql.DB) IItemsRepository {
	return &sItemsRepository{DB: db}
}
