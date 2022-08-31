package main

import (
	"net/http"

	"github.com/Yefhem/hello-world-case/config"
	"github.com/Yefhem/hello-world-case/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	config.ConnectDB()
	
	routes.RoutesInit(r)

	http.ListenAndServe(":8080", r)
}
