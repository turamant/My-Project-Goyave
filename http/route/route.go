package route

import (
	"net/http"
	"goyave.dev/goyave/v4"
)

// Register all the routes
func Register(router *goyave.Router) {
	router.Get("/hello", hello)
}

// Handler function for the "/hello" route
func hello(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "Hi!")
}