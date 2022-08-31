package service

import (
	"errors"

	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/repository"
)

type TicketService interface {
	CreateTicket(tic models.Ticket) (models.Ticket, error)
	GetTicket(id uint) (models.Ticket, error)
	PurchaseTicket(id uint, purchase models.Purchase) error
}

type ticketService struct {
	ticketRepository repository.TicketRepository
}

func NewTicketService(ticketRepo repository.TicketRepository) TicketService {
	return &ticketService{
		ticketRepository: ticketRepo,
	}
}

// --------------------> Methods

// ----------> Create Ticket_option
func (t *ticketService) CreateTicket(tic models.Ticket) (models.Ticket, error) {
	ticket, err := t.ticketRepository.Create(tic)
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

// ----------> Get Ticket by ID
func (t *ticketService) GetTicket(id uint) (models.Ticket, error) {

	ticket, err := t.ticketRepository.Get(id)
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

// ----------> Update Ticket for Purchase
func (t *ticketService) PurchaseTicket(id uint, purchase models.Purchase) error {

	existTicket, err := t.GetTicket(id)
	if err != nil {
		return err
	}

	if purchase.Quantity > existTicket.Allocation {
		return errors.New("more tickets than are available")
	}

	existTicket.Allocation = (existTicket.Allocation - purchase.Quantity)

	if err := t.ticketRepository.Update(existTicket); err != nil {
		return err
	}

	return nil
}
