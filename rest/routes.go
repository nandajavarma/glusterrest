package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/", Index)

	// Volume Life Cycle APIs
	router.HandleFunc("/v1/volumes/{volName}", VolumeCreate).Methods("PUT")
	router.HandleFunc("/v1/volumes/{volName}/start", VolumeStart).Methods("POST")
	router.HandleFunc("/v1/volumes/{volName}/stop", VolumeStop).Methods("POST")
	router.HandleFunc("/v1/volumes/{volName}", VolumeDelete).Methods("DELETE")
	router.HandleFunc("/v1/volumes", VolumeGet).Methods("GET")
	router.HandleFunc("/v1/volumes/{volName}", VolumeGet).Methods("GET")

	// Peers
	// router.HandleFunc("/v1/peers", PeersAttach).Methods("POST")
	// router.HandleFunc("/v1/peers/{fqdn}", PeersDetach).Methods("DELETE")
	// router.HandleFunc("/v1/peers", PeersGet).Methods("GET")
	// router.HandleFunc("/v1/peers/{fqdn}", PeersGet).Methods("GET")

	router.HandleFunc("/v1/events", EventsHandler)
	router.HandleFunc("/v1/listen", ListenHandler).Methods("POST")

	http.Handle("/",
		RestLoggingHandler(
			SetApplicationHeaderJson(
				SignHandler(router))))
}
