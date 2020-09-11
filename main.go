package main
//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=go-boilerplate/templates

import (
	"go-boilerplate/configurations"
	"go-boilerplate/routes"
	"go-boilerplate/utilities"
	"net/http"
)

func main() {
	configuration := configurations.GetConfiguration()

	port := configuration.Port

	utilities.PrintStartMessage(port)

	router := routes.GetRouter()

	panic(http.ListenAndServe(":"+port, router))
}
