package repository

import (
	"github.com/Yefhem/hello-world-case/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	Control(uuid string) (models.User, error)
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

// ----------> Control to User  
func (c *userConnection) Control(uuid string) (models.User, error) {
	var user models.User

	if err := c.connection.First(&user, "id = ?", uuid).Error; err != nil {
		return user, err
	}

	return user, nil
}
