package router

import (
	"F.A.N.A.P/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//routing
	r.HandleFunc("/", controller.GetAllRectangles).Methods("GET")
	r.HandleFunc("/", controller.Insert).Methods("POST")

	return r
}
