package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Yefhem/hello-world-case/config"
	"github.com/Yefhem/hello-world-case/controller"
	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/repository"
	"github.com/Yefhem/hello-world-case/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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

var mockTicketsTrue = []models.Ticket{
	{ID: 1000, Name: "Berlin", Description: "example 1", Allocation: 300},
	{ID: 1001, Name: "Hamilton ", Description: "example 2", Allocation: 400},
	{ID: 1002, Name: "Ankara", Description: "example 3", Allocation: 500},
}

func SetUpRouter() *mux.Router {
	testRouter := mux.NewRouter()

	return testRouter
}

// ----------> This code snippet send some POST request with a sample payload, checks response code and request body with mock data
func TestCreateTicketSuccess(t *testing.T) {

	testRouter := SetUpRouter()

	testRouter.HandleFunc("/ticket_options_test", ticketController.CreateTicket).Methods("POST")

	for _, v := range mockTicketsTrue {

		jsonValue, _ := json.Marshal(v)

		req, err := http.NewRequest("POST", "/ticket_options_test", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}

		newRecorder := httptest.NewRecorder()
		testRouter.ServeHTTP(newRecorder, req)

		data, err := ioutil.ReadAll(newRecorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		var testTickets models.Ticket

		if err := json.Unmarshal(data, &testTickets); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusCreated, newRecorder.Code)
		assert.Equal(t, v, testTickets)
	}

}
func TestGetTicket(t *testing.T) {
	testRouter := SetUpRouter() 

	testRouter.HandleFunc("/ticket_test/{id}", ticketController.GetTicket).Methods("GET")

	for _, v := range mockTicketsTrue {
		id := strconv.FormatUint(v.ID, 10)

		req, err := http.NewRequest("GET", "/ticket_test/"+id, nil)
		if err != nil {
			t.Fatal(err)
		}

		newRecorder := httptest.NewRecorder()
		testRouter.ServeHTTP(newRecorder, req)

		data, _ := ioutil.ReadAll(newRecorder.Body)

		var getTic models.Ticket

		if err := json.Unmarshal(data, &getTic); err != nil {
			t.Fatal(err)
		}

		log.Println(getTic)

		assert.Equal(t, http.StatusOK, newRecorder.Code) // http kodu kontrol et
		assert.EqualValues(t, v, getTic)
	}

}
func TestPurchaseTicket(t *testing.T) {
	testRouter := SetUpRouter() // -- test router oluştur

	// --- 1 tane mock veri post(create) et

	testRouter.HandleFunc("/ticket_options_test", ticketController.CreateTicket).Methods("POST")

	mockTicket := models.Ticket{
		ID:          1011,
		Name:        "konya",
		Description: "example dd",
		Allocation:  50,
	}

	jsonValue, _ := json.Marshal(mockTicket)

	req1, err := http.NewRequest("POST", "/ticket_options_test", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	newRecorder1 := httptest.NewRecorder()
	testRouter.ServeHTTP(newRecorder1, req1)

	assert.Equal(t, http.StatusCreated, newRecorder1.Code)

	// --- bilet satın al

	testRouter.HandleFunc("/ticket_options/{id}/purchases_test", ticketController.PurchaseTicket).Methods("POST")

	mockPurchase := models.Purchase{
		Quantity: 10,
		UserID:   "asadsda2asdas",
	}

	jsonValue2, _ := json.Marshal(mockPurchase)

	id := "1011"

	req2, err := http.NewRequest("POST", "/ticket_options/"+id+"/purchases_test", bytes.NewBuffer(jsonValue2))
	if err != nil {
		t.Fatal(err)
	}

	newRecorder2 := httptest.NewRecorder()
	testRouter.ServeHTTP(newRecorder2, req2)

	result, _ := ticketService.GetTicket(1011)

	log.Println(mockTicket.Allocation - mockPurchase.Quantity)
	log.Println(result.Allocation)

	assert.Equal(t, http.StatusOK, newRecorder2.Code)
	assert.Equal(t, (mockTicket.Allocation - mockPurchase.Quantity), result.Allocation)

}
