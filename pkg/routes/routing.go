package routes

import (
	"web/auth"

	"github.com/gorilla/mux"
)

func Routing() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/", auth.AuthCheck)
	router.HandleFunc("/main", MainPage)
	router.HandleFunc("/create", CreateUser)
	router.HandleFunc("/remove/{id:[0-9]+}", RemoveUser)
	router.HandleFunc("/edit/{id:[0-9]+}", Edit_Get).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", Edit_Post).Methods("POST")

	return router
}