package service

import (
	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/repository"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// --------------------> Methods

// ----------> Create a New User Account
func (u *userService) CreateUser(user models.User) (models.User, error) {

	// ----------> Create User ID with uuid
	user.ID = uuid.NewString()

	result, err := u.userRepository.Create(user)
	if err != nil {
		return result, err
	}
	return result, nil
}
