package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aravindavk/glusterrest/grutil"
	"github.com/gorilla/mux"
)

func main() {
	Autoload()
	router := mux.NewRouter().StrictSlash(true)
	AddRoutes(router)

	port_data := fmt.Sprintf(":%d", grutil.PORT)
	if grutil.HTTPS {
		log.Fatal(http.ListenAndServeTLS(port_data, grutil.CERT, grutil.KEY, nil))
	} else {
		log.Fatal(http.ListenAndServe(port_data, nil))
	}
}
