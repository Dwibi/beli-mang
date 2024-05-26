package merchantrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sMerchantRepository struct {
	DB *sql.DB
}

type IMerchantRepository interface {
	FindMany(filters *entities.SearchMerchantParams) (*entities.MerchantResult, error)
	Create(m *entities.CreateMerchantParams) (int, error)
}

func New(db *sql.DB) IMerchantRepository {
	return &sMerchantRepository{DB: db}
}
