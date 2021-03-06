package service

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// This is the main abstraction file for the this service

// NewServerFromCFEnv decides the url to use for a webClient
func NewServerFromCFEnv(appEnv *cfenv.App) *negroni.Negroni {
	webClient := fulfillmentWebClient{
		rootURL: "http://localhost:3001/skus",
	}

	val, err := cftools.GetVCAPServiceProperty("backing-fulfillment", "url", appEnv)
	if err != nil {
		webClient.rootURL = val
	} else {
		fmt.Printf("Failed to get URL property from bound service: %v \n", err)
	}

	fmt.Printf("Using the following URL for fulfilment backing service: %s \n", webClient.rootURL)

	return NewServerFromClient(webClient)
}

// NewServer generates a real server that has access to the address of a
// dependent service (fulfilment) to make requests
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	webClient := fulfillmentWebClient{
		rootURL: "http://localhost:3001/skus/",
	}

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)

	return n

}

// NewServerFromClient configure and retuns a server
func NewServerFromClient(webClient fulfillmentClient) *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)

	return n

}

// NewServerFromClient configures with a fake backing service (fulfilment service)
// call and returns a Server that was generated by the client
// func NewServerFromClient(webClient fulfillmentClient) *negroni.Negroni {
// 	formatter := render.New(render.Options{
// 		IndentJSON: true,
// 	})
//
// 	n := negroni.Classic()
// 	mx := mux.NewRouter()
//
// 	initRoutes(mx, formatter, webClient)
//
// 	n.UseHandler(mx)
//
// 	return n
//
// }

func initRoutes(mx *mux.Router, formatter *render.Render, webClient fulfillmentClient) {
	mx.HandleFunc("/", rootHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog", getAllCatalogItemsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog/{sku}", getCatalogItemDetailsHander(formatter, webClient)).Methods("GET")
}
