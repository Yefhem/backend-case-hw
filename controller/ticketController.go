package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Yefhem/hello-world-case/helpers"
	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/service"
	"github.com/gorilla/mux"
)

type TicketController interface {
	CreateTicket(w http.ResponseWriter, r *http.Request)
	GetTicket(w http.ResponseWriter, r *http.Request)
	PurchaseTicket(w http.ResponseWriter, r *http.Request)
}

type ticketController struct {
	ticketService service.TicketService
	userService   service.UserService
}

func NewTicketController(ticketService service.TicketService, userService service.UserService) TicketController {
	return &ticketController{
		ticketService: ticketService,
		userService:   userService,
	}
}

// --------------------> Methods

func (cont *ticketController) CreateTicket(w http.ResponseWriter, r *http.Request) {

	var ticket models.Ticket

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	if err := json.Unmarshal(body, &ticket); err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	if helpers.IsItEmpty(ticket.Name) || helpers.IsItEmpty(ticket.Description) || (ticket.Allocation == 0) {
		helpers.Response(w, http.StatusBadRequest, "bad request")
		return
	}

	ticketModel, err := cont.ticketService.CreateTicket(ticket)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.Response(w, http.StatusCreated, ticketModel)
}

func (cont *ticketController) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	ticket, err := cont.ticketService.GetTicket(uint(id))
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.Response(w, http.StatusOK, ticket)
}

func (cont *ticketController) PurchaseTicket(w http.ResponseWriter, r *http.Request) {

	var purchase models.Purchase

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	username, pass, ok := r.BasicAuth()
	if !ok {
		helpers.Response(w, http.StatusUnauthorized, nil)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	if err := json.Unmarshal(body, &purchase); err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	if err := cont.userService.ControlAuth(purchase.UserID, username, pass); err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusUnauthorized, err.Error())
		return
	}

	if err := cont.ticketService.PurchaseTicket(uint(id), purchase); err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.Response(w, http.StatusOK, "")
}
