package routes

import (
	"github.com/gorilla/mux"
)

func Routing() *mux.Ser{
	router := mux.NewRouter()
	router.HandleFunc("/add", CreateUser())
	router.HandleFunc("/remove", RemoveUser())
	router.HandleFunc("/edit/{id:[0-9]+}", Edit_Get()).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", Edit_Post()).Methods("POST")

	return router
}