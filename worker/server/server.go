package server

import "github.com/gorilla/mux"

type Options struct {
	ListenAddr string
}

type server struct {
	opt    Options
	router *mux.Router
}
