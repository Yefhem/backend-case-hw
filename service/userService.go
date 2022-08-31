package service

import (
	"errors"
	"log"

	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	ControlAuth(purchaseID, username, pass string) error
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

	user.Password = u.PasswordHasher(user.Password)

	result, err := u.userRepository.Create(user)
	if err != nil {
		return result, err
	}
	return result, nil
}

// ----------> Control Authentication (Basic Auth)
func (u *userService) ControlAuth(purchaseID, username, pass string) error {
	user, err := u.userRepository.Control(purchaseID)
	if err != nil {
		return err
	}

	storedUserPass := user.Password

	// if user.Username != username || user.Password != pass {
	// 	return errors.New("unauthorized")
	// }

	if user.Username != username || !(u.VerifyPassword(storedUserPass, pass))  {
		return errors.New("unauthorized")
	}

	return nil
}

// ----------> Hash to User Password
func (u *userService) PasswordHasher(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// ----------> Verify to User Hashed Password
func (u *userService) VerifyPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		log.Println(err)
		return false
	}

	return true
}
