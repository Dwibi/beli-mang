package userrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type IUserRepository interface {
	IsExist(*IsExistParams) (bool, error)
	Create(*ParamsCreateUser) (int, error)
	FindOne(*FindOneParams) (*entities.User, error)
}

func New(db *sql.DB) IUserRepository {
	return &sUserRepository{DB: db}
}
