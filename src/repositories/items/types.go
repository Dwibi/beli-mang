package itemsrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sItemsRepository struct {
	DB *sql.DB
}

type IItemsRepository interface {
	FindMany(merchantId int, filters *entities.SearchItemsParams) (*entities.ItemsResult, error)
	Create(merchantId int, m *entities.CreateItemsParams) (int, error)
	GetMissingIDs(ids []int) ([]int, error)
}

func New(db *sql.DB) IItemsRepository {
	return &sItemsRepository{DB: db}
}
