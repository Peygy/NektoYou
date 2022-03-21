package server

import (
	"net/http"
	"web/pkg/routes"
)

func Launch(port string) error{
	http.Handle("/", routes.Routing())
	err := http.ListenAndServe(port, nil)
	return err
}