package userusecase

import (
	"github.com/Dwibi/beli-mang/src/entities"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
)

type sUserUseCase struct {
	userRepository userrepository.IUserRepository
}

type IUserUseCase interface {
	AdminRegister(entities.RegisterParams) (string, int, error)
	UserRegister(entities.RegisterParams) (string, int, error)
	AdminLogin(entities.LoginParams) (string, int, error)
	UserLogin(entities.LoginParams) (string, int, error)
}

func New(userRepository userrepository.IUserRepository) IUserUseCase {
	return &sUserUseCase{
		userRepository: userRepository,
	}
}
