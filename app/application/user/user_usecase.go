package user

import "github.com/Timming/app/domain"

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) *userUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}