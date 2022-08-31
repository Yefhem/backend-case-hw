package routes

import (
	"github.com/Yefhem/hello-world-case/config"
	"github.com/Yefhem/hello-world-case/controller"
	"github.com/Yefhem/hello-world-case/repository"
	"github.com/Yefhem/hello-world-case/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	// ----------> db
	db *gorm.DB = config.ConnectDB()
	// ----------> repository layer
	ticketRepository repository.TicketRepository = repository.NewTicketRepository(db)
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	// ----------> service layer
	ticketService service.TicketService = service.NewTicketService(ticketRepository)
	userService   service.UserService   = service.NewUserService(userRepository)
	// ----------> controller layer
	ticketController controller.TicketController = controller.NewTicketController(ticketService, userService)
	userController   controller.UserController   = controller.NewUserController(userService)
)

func RoutesInit(router *mux.Router) {
	// ----------> Ticket
	router.HandleFunc("/ticket_options", ticketController.CreateTicket).Methods("POST")
	router.HandleFunc("/ticket_options/{id}/purchases", ticketController.PurchaseTicket).Methods("POST")
	router.HandleFunc("/ticket/{id}", ticketController.GetTicket).Methods("GET")

	// ----------> User
	router.HandleFunc("/user_create", userController.CreateUser).Methods("POST")

}
