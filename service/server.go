package main

import (
	"net/http"

	"github.com/benschw/opin-go/ophttp"
	"github.com/benschw/opin-go/rest"
)

// Resource Handler for `/greeting`
func GreetingHandler(resp http.ResponseWriter, req *http.Request) {
	rest.SetOKResponse(resp, "hello world")
}

// Wire and start http server
func RunServer(server *ophttp.Server) {
	http.Handle("/greeting", http.HandlerFunc(GreetingHandler))
	server.Start()
}
