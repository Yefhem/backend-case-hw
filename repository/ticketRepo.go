package repository

import (
	"github.com/Yefhem/hello-world-case/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket models.Ticket) (models.Ticket, error)
	Get(id uint) (models.Ticket, error)
	Update(updatedTicket models.Ticket) error
}

type ticketConnection struct {
	connection *gorm.DB
}

func NewTicketRepository(connection *gorm.DB) TicketRepository {
	return &ticketConnection{
		connection: connection,
	}
}

// --------------------> Methods

// ----------> Create to DB
func (c *ticketConnection) Create(ticket models.Ticket) (models.Ticket, error) {
	if err := c.connection.Create(&ticket).Error; err != nil {
		return ticket, err
	}
	return ticket, nil
}

// ----------> Get Single Data from DB by ID
func (c *ticketConnection) Get(id uint) (models.Ticket, error) {
	var ticket models.Ticket

	if err := c.connection.First(&ticket, "id = ?", id).Error; err != nil {
		return ticket, err
	}

	return ticket, nil
}

// ----------> Save Changes to DB
func (c *ticketConnection) Update(updatedTicket models.Ticket) error {

	if err := c.connection.Save(&updatedTicket).Error; err != nil {
		return err
	}

	return nil
}
