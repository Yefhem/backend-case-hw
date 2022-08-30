package repository

import (
	"github.com/Yefhem/hello-world-case/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userConnection{
		connection: connection,
	}
}

// --------------------> Methods

// ----------> Create a New User Account
func (c *userConnection) Create(user models.User) (models.User, error) {

	user.ID = uuid.NewString()

	if err := c.connection.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
