package repository

import (
	"github.com/Yefhem/hello-world-case/models"
	"gorm.io/gorm"
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

// ----------> Create to DB
func (c *userConnection) Create(user models.User) (models.User, error) {
	if err := c.connection.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
